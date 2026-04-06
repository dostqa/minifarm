package storage

import (
	"image"
	_ "image/png"
	"log"
	"minifarm/internal/gametypes"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// state - описывает в каком состоянии находится сущность.
// Если сущность соверщает движение, idle = false
type state struct {
	isIdle bool
	facing gametypes.Vector
}

var (
	MovesUp    = state{false, gametypes.UpVector}
	MovesRight = state{false, gametypes.RightVector}
	MovesDown  = state{false, gametypes.DownVector}
	MovesLeft  = state{false, gametypes.LeftVector}
	LooksUp    = state{true, gametypes.UpVector}
	LooksRight = state{true, gametypes.RightVector}
	LooksDown  = state{true, gametypes.DownVector}
	LooksLeft  = state{true, gametypes.LeftVector}
)

var spriteViews = map[state]string{
	MovesUp:    "/moves_up.png",
	LooksUp:    "/looks_up.png",
	MovesRight: "/moves_right.png",
	LooksRight: "/looks_right.png",
	MovesDown:  "/moves_down.png",
	LooksDown:  "/looks_down.png",
	MovesLeft:  "/moves_left.png",
	LooksLeft:  "/looks_left.png",
}

var DefaultAssetStorage *AssetStorage

func init() {
	DefaultAssetStorage = &AssetStorage{
		cache: make(map[string][]*ebiten.Image),
	}
}

// ticker - подсказывает только, какой сейчас тик.
// Каждую секунду начинает отсчет с начала (то есть с нуля)
// и считает до TicksPerSecond() - 1
type ticker interface {
	NowTick() int
	TicksPerSecond() int
}

// SpriteInfoProvider предоставляет информацию о спрайте и его анимации
type SpriteInfoProvider interface {
	ID() string
	FrameCount() int
	FrameWidth() int
	FrameHeight() int
}

// SpriteStateInfoProvider предоставляет информацию о своём состоянии
type SpriteStateInfoProvider interface {
	IsIdle() bool
	Facing() gametypes.Vector
}

type AssetStorage struct {
	ticker ticker
	// В кэш изображения складываются только для одного состояния (одного файлового пути)
	cache map[string][]*ebiten.Image
}

func (storage *AssetStorage) GetSprite(provider SpriteInfoProvider) *ebiten.Image {
	nowFrame := storage.nowFrame(provider.FrameCount())
	path := storage.getPath(provider)

	// Пытаемся получить спрайт из кэша
	sprite, ok := storage.getSpriteFromCacheByPath(path, nowFrame)
	if ok {
		return sprite
	}

	// Если в кэше нет нужного спрайта:
	// Загружаем спрайтшит с диска
	spriteSheet := storage.loadSpriteSheetByPath(path)
	// Нарезаем спрайтшит на кадры
	frames := storage.sliceSpriteSheet(spriteSheet, provider.FrameWidth(), provider.FrameHeight(), provider.FrameCount())
	// Сохраняем кадры в кэш
	storage.cache[path] = frames

	return frames[nowFrame]
}

func (storage *AssetStorage) getPath(provider SpriteInfoProvider) string {
	path := "./assets/sprites/" + provider.ID() + "/animation.png"

	if s, ok := provider.(SpriteStateInfoProvider); ok {
		state := state{s.IsIdle(), s.Facing()}
		path = "./assets/sprites/" + provider.ID() + spriteViews[state]
	}
	return path
}

func (storage *AssetStorage) getSpriteFromCacheByPath(path string, nowFrame int) (*ebiten.Image, bool) {
	if frames, ok := storage.cache[path]; ok {
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
