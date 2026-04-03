package storage

import (
	"image"
	_ "image/png"
	"log"
	"minifarm/internal/gametypes"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FrameWidth  = 16
	FrameHeight = 16
	FrameCount  = 4
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

// ticker - подсказывает только, какой сейчас кадр.
// Каждую секунду начинает отсчет с начала (то есть с нуля)
type ticker interface {
	NowFrame() int
}

type AssetStorage struct {
	ticker ticker
	// В кэш изображения складываются только для одного состояния (одного файлового пути)
	cache map[string][]*ebiten.Image
}

func (storage *AssetStorage) GetDirectionalSprite(id string, isIdle bool, facing gametypes.Vector) *ebiten.Image {
	state := state{isIdle, facing}
	path := "./assets/sprites/" + id + spriteViews[state]
	return storage.getSpriteByPath(path)
}

func (storage *AssetStorage) GetSingleSprite(id string) *ebiten.Image {
	path := "./assets/sprites/" + id + "/animation"
	return storage.getSpriteByPath(path)
}

func (storage *AssetStorage) getSpriteByPath(path string) *ebiten.Image {
	if frames, ok := storage.cache[path]; ok {
		return frames[storage.ticker.NowFrame()]
	}

	img := storage.loadSpriteByPath(path)
	frames := storage.sliceSpriteSheet(img)
	storage.cache[path] = frames

	return frames[storage.ticker.NowFrame()]
}

func (storage *AssetStorage) loadSpriteByPath(path string) *ebiten.Image {
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

func (storage *AssetStorage) sliceSpriteSheet(img *ebiten.Image) []*ebiten.Image {
	frames := make([]*ebiten.Image, FrameCount)

	for i := 0; i < FrameCount; i++ {
		frame := image.Rect(
			i*FrameWidth, 0, // левый верхний угол
			(i+1)*FrameWidth, FrameHeight, // правый нижний угол
		)
		frames[i] = ebiten.NewImageFromImage(img.SubImage(frame))
	}

	return frames
}

func (storage *AssetStorage) ConnectToTicker(ticker ticker) {
	storage.ticker = ticker
}
