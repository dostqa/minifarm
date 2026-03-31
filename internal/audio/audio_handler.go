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
	DefaultAudioHandler *AudioHandler
)

func init() {
	AudioGameContext = audio.NewContext(sampleRate)
	DefaultAudioHandler = &AudioHandler{context: AudioGameContext, cache: make(map[string][]byte)}
}

type AudioHandler struct {
	context *audio.Context
	cache   map[string][]byte
}

// TODO:
// Обрабатывать ошибки, а не вызывать log.Fatal(err)
// Думаю в кэше надо хранить Player'ов

func (handler *AudioHandler) LoadStreamFromMP3(path string) *mp3.Stream {
	// Проверка: нет ли уже данных в кэше
	if data, ok := handler.cache[path]; ok {
		return handler.LoadStreamFromData(data)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	handler.cache[path] = data // Добавляем данные в кэш

	return handler.LoadStreamFromData(data)
}

func (handler *AudioHandler) LoadStreamFromData(data []byte) *mp3.Stream {
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
		stream := handler.LoadStreamFromMP3("./assets/audio/axe.mp3")
		player, err := handler.context.NewPlayer(stream)
		if err != nil {
			log.Fatal(err)
		}

		player.Play()
	}
}

// Напоминалка самому себе:
// Таким образом подготавливаем файл, если он большой: (файл частями попадает в память)
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
