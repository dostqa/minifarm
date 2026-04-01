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

var DefaultAssetStorage *AssetStorage

func init() {
	DefaultAssetStorage = &AssetStorage{
		cache: make(map[string][]*ebiten.Image),
	}
}

// ticker - подсказывает только, какой сейчас кадр.
// Каждую секунду начинает отсчет сначала
type ticker interface {
	NowFrame() int
}

// Condition - описывает в каком состоянии находится сущность.
// Если сущность соверщает движение velocity != {0, 0}
type Condition struct {
	velocity gametypes.Vector
	facing   gametypes.Vector
}

var (
	MovesUp    = Condition{gametypes.UpVector, gametypes.UpVector}
	MovesRight = Condition{gametypes.RightVector, gametypes.RightVector}
	MovesDown  = Condition{gametypes.DownVector, gametypes.DownVector}
	MovesLeft  = Condition{gametypes.LeftVector, gametypes.LeftVector}
	LooksUp    = Condition{gametypes.ZeroVector, gametypes.UpVector}
	LooksRight = Condition{gametypes.ZeroVector, gametypes.RightVector}
	LooksDown  = Condition{gametypes.ZeroVector, gametypes.DownVector}
	LooksLeft  = Condition{gametypes.ZeroVector, gametypes.LeftVector}
)

func NewCondition(velocity, facing gametypes.Vector) Condition {
	return Condition{velocity: velocity, facing: facing}
}

type SpritesID = map[Condition]string

var PlayerSprites = SpritesID{
	MovesUp:    "./assets/sprites/player_moves_up.png",
	MovesRight: "./assets/sprites/player_moves_right.png",
	MovesDown:  "./assets/sprites/player_moves_down.png",
	MovesLeft:  "./assets/sprites/player_moves_Left.png",
	LooksUp:    "./assets/sprites/player_looks_up.png",
	LooksRight: "./assets/sprites/player_looks_right.png",
	LooksDown:  "./assets/sprites/player_looks_down.png",
	LooksLeft:  "./assets/sprites/player_looks_left.png",
}

type AssetStorage struct {
	ticker ticker
	// В кэш изображения складываются только для одного состояния (одного пути)
	cache map[string][]*ebiten.Image
}

func (storage *AssetStorage) GetSpriteByName(path string) *ebiten.Image {
	if frames, ok := storage.cache[path]; ok {
		return frames[storage.ticker.NowFrame()]
	}

	img := storage.loadSpriteByName(path)
	frames := storage.sliceSpriteSheet(img)
	storage.cache[path] = frames

	return frames[storage.ticker.NowFrame()]
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

func (storage *AssetStorage) loadSpriteByName(path string) *ebiten.Image {
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

func (storage *AssetStorage) ConnectToTicker(ticker ticker) {
	storage.ticker = ticker
}
