package audio

import (
	"bytes"
	"log"
	"minifarm/internal/events"
	"minifarm/internal/gametypes"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

var (
	sampleRate          = 48000        // частота дискретизации
	AudioGameContext    *audio.Context // Централизованное хранилище состояний аудиосистемы
	DefaultAudioHandler *AudioHandler  // DefaultAudioHandler реагируют на события, запуская звуки
)

func init() {
	AudioGameContext = audio.NewContext(sampleRate)
	DefaultAudioHandler = &AudioHandler{context: AudioGameContext, cache: make(map[SoundID]*audio.Player)}
}

// SoundID используется для идентификации звука
type SoundID = string

const (
	Axe  SoundID = "./assets/audio/axe.mp3"
	Step SoundID = "./assets/audio/step.mp3"
)

type AudioHandler struct {
	context *audio.Context
	cache   map[SoundID]*audio.Player // в кэше хранятся Player'ы

	// Хранятся именно Player'ы, а не сырые данные,
	// Чтобы не наслаивать звуки друг на друга
}

// TODO:
// Обрабатывать ошибки, а не вызывать log.Fatal(err)

func (handler *AudioHandler) loadStreamFromMP3(id SoundID) *mp3.Stream {
	// читаем файл и помещаем весь файл в память
	data, err := os.ReadFile(id)
	if err != nil {
		log.Fatal(err)
	}

	return handler.loadStreamFromData(data)
}

func (handler *AudioHandler) loadStreamFromData(data []byte) *mp3.Stream {
	reader := bytes.NewReader(data)
	stream, err := mp3.DecodeWithSampleRate(handler.context.SampleRate(), reader)
	if err != nil {
		log.Fatal(err)
	}

	return stream
}

func (handler *AudioHandler) Handle(event events.Event) {
	switch event.Type() {
	case gametypes.ToolUsedEventType:
		if player, err := handler.cache[Axe]; err {
			if player.IsPlaying() {
				return
			} else {
				player.Play()
			}
		}

		player := handler.newPlayerFromMP3(Axe)
		handler.cache[Axe] = player
		player.Play()

	case gametypes.EntityMovedEventType:
		if player, err := handler.cache[Step]; err {
			if player.IsPlaying() {
				return
			} else {
				player.Play()
			}
		}

		player := handler.newPlayerFromMP3(Step)
		handler.cache[Step] = player
		player.Play()
	}
}

func (handler *AudioHandler) newPlayerFromMP3(id SoundID) *audio.Player {
	stream := handler.loadStreamFromMP3(id)

	player, err := handler.context.NewPlayer(stream)
	if err != nil {
		log.Fatal(err)
	}

	return player
}

// Напоминалка самому себе:
// Таким образом подготавливаем файл, если он большой: (файл постепенно частями попадает в память)
// f, err := os.Open("sound.mp3")
// if err != nil {
//     log.Fatal(err)
// }
// defer f.Close()
// var r io.Reader = f

// Таким образом, если не большой: (файл сразу весь попадает в память)
// data, err := os.ReadFile("sound.mp3")
// if err != nil {
//     log.Fatal(err)
// }
// r := bytes.NewReader(data)
