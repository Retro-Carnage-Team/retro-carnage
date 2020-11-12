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
	effects map[assets.SoundEffect]sound
	mixer   *beep.Mixer
}

type sound interface {
	play(mixer *beep.Mixer)
	stop()
}

// basicSound is a sound that gets played exactly one time. No further interaction is possible
type basicSound struct {
	buffer *beep.Buffer
}

func (bs *basicSound) play(mixer *beep.Mixer) {
	mixer.Add(bs.buffer.Streamer(0, bs.buffer.Len()))
}

func (bs *basicSound) stop() {}

// loopingEffect is a sound that gets played over and over again. You can pause & continue
type loopingEffect struct {
	control *beep.Ctrl
}

func (bs *loopingEffect) play(mixer *beep.Mixer) {
	speaker.Lock()
	bs.control.Paused = false
	speaker.Unlock()
}

func (bs *loopingEffect) stop() {
	speaker.Lock()
	bs.control.Paused = true
	speaker.Unlock()
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
	var aSound = sb.effects[effect]
	if nil != aSound {
		aSound.play(sb.mixer)
	}
}

func (sb *SoundBoard) Stop(effect assets.SoundEffect) {
	var aSound = sb.effects[effect]
	if nil != aSound {
		aSound.stop()
	}
}

func (sb *SoundBoard) initialize() {
	err := speaker.Init(mp3Format.SampleRate, mp3Format.SampleRate.N(time.Second/10))
	if err != nil {
		Error.Println(err.Error())
	}

	sb.mixer = &beep.Mixer{}
	speaker.Play(sb.mixer)

	sb.effects = make(map[assets.SoundEffect]sound)
	for _, fx := range assets.SoundEffects {
		buffer, err := bufferMp3File(fx, sb.mixer)
		if err != nil {
			Error.Panicln(err.Error())
		} else {
			sb.effects[fx] = buffer
		}
	}
}

func bufferMp3File(fx assets.SoundEffect, mixer *beep.Mixer) (sound, error) {
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

	if isLoopingEffect(fx) {
		ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, buffer.Streamer(0, buffer.Len())), Paused: true}
		mixer.Add(ctrl)
		return &loopingEffect{control: ctrl}, nil
	}
	return &basicSound{buffer: buffer}, nil
}

func isLoopingEffect(fx assets.SoundEffect) bool {
	for _, v := range assets.LoopingSoundEffects {
		if v == fx {
			return true
		}
	}
	return false
}
