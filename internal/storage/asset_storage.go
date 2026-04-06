package storage

import (
	"encoding/json"
	"image"
	_ "image/png"
	"log"
	"minifarm/internal/gametypes"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

var DefaultAssetStorage *AssetStorage

func init() {
	DefaultAssetStorage = &AssetStorage{
		imageCache: make(map[string][]*ebiten.Image),
		infoCache:  make(map[string]*spriteSheetInfo),
	}
}

func infoPath(id string) string {
	return filepath.Join("assets", "sprites", id, "info.json")
}

func imagePath(id string, state gametypes.StateName, facing gametypes.Vector) string {
	return filepath.Join("assets", "sprites", id, stateToString[state]+facingToString[facing]+".png")
}

type spriteSheetInfo struct {
	FrameCount  int
	FrameWidth  int
	FrameHeight int
}

var stateToString = map[gametypes.StateName]string{
	gametypes.IdleStateName: "idle",
	gametypes.MoveStateName: "moves",
}

var facingToString = map[gametypes.Vector]string{
	gametypes.ZeroVector:  "",
	gametypes.UpVector:    "up",
	gametypes.RightVector: "right",
	gametypes.DownVector:  "down",
	gametypes.LeftVector:  "left",
}

// ticker - подсказывает только, какой сейчас тик.
// Каждую секунду начинает отсчет с начала (то есть с нуля)
// и считает до TicksPerSecond() - 1
type ticker interface {
	NowTick() int
	TicksPerSecond() int
}

type AssetStorage struct {
	ticker ticker
	// В кэш изображения складываются только для одного состояния (одного файлового пути)
	imageCache map[string][]*ebiten.Image
	infoCache  map[string]*spriteSheetInfo
}

func (storage *AssetStorage) GetSprite(id string, state gametypes.StateName, facing gametypes.Vector) *ebiten.Image {
	info := storage.getInfo(infoPath(id))

	// Считаем какой нужен кадр анимации
	nowFrame := storage.nowFrame(info.FrameCount)
	path := imagePath(id, state, facing)

	// Пытаемся получить кадр из кэша
	sprite, ok := storage.getFrameFromCacheByPath(path, nowFrame)
	if ok {
		return sprite
	}

	// Если в кэше нет нужного спрайтшита:
	// Загружаем спрайтшит с диска
	spriteSheet := storage.loadSpriteSheetByPath(path)
	// Нарезаем спрайтшит на кадры
	frames := storage.sliceSpriteSheet(spriteSheet, info.FrameWidth, info.FrameHeight, info.FrameCount)
	// Сохраняем кадры в кэш
	storage.imageCache[path] = frames
	return frames[nowFrame]
}

func (storage *AssetStorage) getInfo(path string) *spriteSheetInfo {
	if info, ok := storage.infoCache[path]; ok {
		return info
	}

	info, err := storage.loadInfo(path)
	if err != nil {
		log.Fatalf("failed to load info: %v", err)
	}
	storage.infoCache[path] = info
	return info
}

func (storage *AssetStorage) loadInfo(path string) (*spriteSheetInfo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	info := &spriteSheetInfo{}
	err = json.Unmarshal(data, info)
	return info, err
}

func (storage *AssetStorage) getFrameFromCacheByPath(path string, nowFrame int) (*ebiten.Image, bool) {
	if frames, ok := storage.imageCache[path]; ok {
		return frames[nowFrame], ok
	}

	return nil, false
}

func (storage *AssetStorage) loadSpriteSheetByPath(path string) *ebiten.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}

	return ebiten.NewImageFromImage(img)
}

func (storage *AssetStorage) sliceSpriteSheet(img *ebiten.Image, frameWidth, frameHeight, frameCount int) []*ebiten.Image {
	frames := make([]*ebiten.Image, frameCount)

	for i := 0; i < frameCount; i++ {
		frame := image.Rect(
			i*frameWidth, 0, // левый верхний угол
			(i+1)*frameWidth, frameHeight, // правый нижний угол
		)
		frames[i] = ebiten.NewImageFromImage(img.SubImage(frame))
	}

	return frames
}

func (storage *AssetStorage) nowFrame(frameCount int) int {
	return storage.ticker.NowTick() / (storage.ticker.TicksPerSecond() / frameCount)
}

func (storage *AssetStorage) ConnectToTicker(ticker ticker) {
	storage.ticker = ticker
}
