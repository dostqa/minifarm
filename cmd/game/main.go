package main

import (
	"fmt"
	"log"
	"minifarm/internal/audio"
	"minifarm/internal/commands"
	"minifarm/internal/events"
	"minifarm/internal/input"
	"minifarm/internal/physic"
	"minifarm/internal/render"
	"minifarm/internal/storage"
	"minifarm/internal/ticker"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	f, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть лог файл: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Программа запущена")

	events.DefaultBus.Subscribe(physic.DefaultHandler, audio.DefaultAudioHandler)

	input.DefaultInput.ConnectToInvoker(&commands.DefaultInvoker)
	storage.DefaultAssetStorage.ConnectToTicker(ticker.DefaultTicker)

	ebiten.SetWindowSize(render.Width*render.ScaleValue, render.Height*render.ScaleValue)
	game := NewGame()

	fmt.Println("Game is started!")
	_ = ebiten.RunGame(game)
}
