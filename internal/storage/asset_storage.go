package storage

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"minifarm/internal/ticker"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FrameWidth  = 16
	FrameHeight = 16
	FrameCount  = 4
)

var DefaultAssetStorage *AssetStorage

func init() {
	DefaultAssetStorage = &AssetStorage{
		ticker: ticker.DefaultTicker,
		cache:  make(map[SpriteID][]*ebiten.Image),
	}
}

type SpriteID string

const (
	PlayerSprite SpriteID = "./assets/sprites/player.png"
)

type AssetStorage struct {
	ticker *ticker.Ticker
	cache  map[SpriteID][]*ebiten.Image
}

func (storage *AssetStorage) GetSpriteByID(path SpriteID) *ebiten.Image {
	if frames, ok := storage.cache[path]; ok {
		return frames[ticker.DefaultTicker.NowFrame()]
	}

	img := storage.loadSpriteByName(path)
	frames := storage.sliceSpriteSheet(img)
	storage.cache[path] = frames

	return frames[ticker.DefaultTicker.NowFrame()]
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

func (storage *AssetStorage) loadSpriteByName(path SpriteID) *ebiten.Image {
	fmt.Println(string(path))
	file, err := os.Open(string(path))
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
