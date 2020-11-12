package util

import (
	"errors"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"os"
	"path/filepath"
	"retro-carnage.net/assets"
	"time"
)

type SoundBoard struct {
	effects map[assets.SoundEffect]*decodedSoundFile
}

type sound interface {
	play()
	stop()
}

// basicSound is a sound that gets played exactly one time. No further interaction is possible
type basicSound struct {
	soundFile decodedSoundFile
}

func (bs *basicSound) play() {

}

func (bs *basicSound) stop() {}

type decodedSoundFile struct {
	buffer *beep.Buffer
	format beep.Format
}

// This is the format we use for all music assets (music and effects). There is no need to analyze each file.
var mp3Format = beep.Format{SampleRate: 32000, NumChannels: 2, Precision: 2}

var soundboard *SoundBoard

func NewSoundBoard() *SoundBoard {
	if nil == soundboard {
		soundboard = &SoundBoard{}
		soundboard.initialize()
	}
	return soundboard
}

func (sb *SoundBoard) Play(effect assets.SoundEffect) {
	var decodedSoundFile = sb.effects[effect]
	if nil != decodedSoundFile {
		sound := decodedSoundFile.buffer.Streamer(0, decodedSoundFile.buffer.Len())
		speaker.Play(sound)
	}
}

func (sb *SoundBoard) initialize() {
	sb.effects = make(map[assets.SoundEffect]*decodedSoundFile)
	for _, fx := range assets.SoundEffects {
		buffer, err := bufferMp3File(fx)
		if err != nil {
			Error.Panicln(err.Error())
		} else {
			sb.effects[fx] = buffer
		}
	}

	speaker.Init(mp3Format.SampleRate, mp3Format.SampleRate.N(time.Second/10))
}

func bufferMp3File(fx assets.SoundEffect) (*decodedSoundFile, error) {
	var filePath = filepath.Join(".", "sounds", "fx", string(fx))
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to load sound effect from %s: %v", filePath, err))
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to decode sound effect from %s: %v", filePath, err))
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	_ = streamer.Close()
	return &decodedSoundFile{buffer: buffer, format: format}, nil
}

func isLoopingEffect(fx assets.SoundEffect) bool {
	for _, v := range assets.LoopingSoundEffects {
		if v == fx {
			return true
		}
	}
	return false
}
