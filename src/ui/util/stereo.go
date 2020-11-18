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
	"retro-carnage.net/logging"
	"retro-carnage.net/util"
	"time"
)

// Stereo is the class we use to play music and sound effects throughout the application.
type Stereo struct {
	effects map[assets.SoundEffect]sound
	mixer   *beep.Mixer
	music   map[assets.Song]sound
}

var stereo *Stereo

type sound interface {
	play(mixer *beep.Mixer)
	stop()
}

//--- Stereo ---------------------------------------------------------------------------------------------------------//

// NewStereo initializes and returns the Singleton instance of Stereo
func NewStereo() *Stereo {
	if nil == stereo {
		stereo = &Stereo{}
		stereo.initialize()
	}
	return stereo
}

func (sb *Stereo) initialize() {
	var mp3Format = beep.Format{SampleRate: 32000, NumChannels: 2, Precision: 2}
	err := speaker.Init(mp3Format.SampleRate, mp3Format.SampleRate.N(time.Second/10))
	if err != nil {
		logging.Error.Println(err.Error())
	}

	sb.mixer = &beep.Mixer{}
	speaker.Play(sb.mixer)

	sb.effects = make(map[assets.SoundEffect]sound)
	for _, fx := range assets.SoundEffects {
		sound, err := loadSoundEffect(fx)
		if err != nil {
			logging.Error.Panicln(err.Error())
		} else {
			sb.effects[fx] = sound
		}
	}

	sb.music = make(map[assets.Song]sound)
}

// PlayFx starts the playback of a given SoundEffect
func (sb *Stereo) PlayFx(effect assets.SoundEffect) {
	var aSound = sb.effects[effect]
	if nil != aSound {
		aSound.play(sb.mixer)
	}
}

// StopFX immediately stops the playback of a given SoundEffect
func (sb *Stereo) StopFx(effect assets.SoundEffect) {
	var aSound = sb.effects[effect]
	if nil != aSound {
		aSound.stop()
	}
}

// PlaySong starts the playback of a given Song
func (sb *Stereo) PlaySong(song assets.Song) {
	var aSound = sb.music[song]
	if nil == aSound {
		var err error = nil
		aSound, err = loadMusic(song)
		if err != nil {
			logging.Error.Panicln(err.Error())
		}
		sb.music[song] = aSound
	}
	aSound.play(sb.mixer)
}

// PlaySong immediately stops the playback of a given Song
func (sb *Stereo) StopSong(song assets.Song) {
	var aSound = sb.music[song]
	if nil != aSound {
		aSound.stop()
	}
}

func loadSoundEffect(fx assets.SoundEffect) (sound, error) {
	stopWatch := util.StopWatch{Name: "Buffering sound effect " + string(fx)}
	stopWatch.Start()

	var filePath = filepath.Join(".", "sounds", "fx", string(fx))
	buffer, err := readMp3IntoBuffer(filePath)
	if err != nil {
		return nil, err
	}

	_ = stopWatch.Stop()
	logging.Trace.Println(stopWatch.DebugMessage())

	if isLoopingEffect(fx) {
		return &loopingSound{buffer: buffer}, nil
	}
	return &basicSound{buffer: buffer}, nil
}

func loadMusic(song assets.Song) (sound, error) {
	stopWatch := util.StopWatch{Name: "Buffering music: " + string(song)}
	stopWatch.Start()

	var filePath = filepath.Join(".", "sounds", "music", string(song))
	buffer, err := readMp3IntoBuffer(filePath)
	if err != nil {
		return nil, err
	}

	_ = stopWatch.Stop()
	logging.Trace.Println(stopWatch.DebugMessage())

	return &loopingSound{buffer: buffer}, nil
}

func readMp3IntoBuffer(filePath string) (*beep.Buffer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to load music from %s: %v", filePath, err))
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to decode music from %s: %v", filePath, err))
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	_ = streamer.Close()

	return buffer, nil
}

func isLoopingEffect(fx assets.SoundEffect) bool {
	for _, v := range assets.LoopingSoundEffects {
		if v == fx {
			return true
		}
	}
	return false
}

//--- Helpers --------------------------------------------------------------------------------------------------------//

// basicSound is a sound that gets played exactly one time. No further interaction is possible
type basicSound struct {
	buffer *beep.Buffer
}

func (bs *basicSound) play(mixer *beep.Mixer) {
	mixer.Add(bs.buffer.Streamer(0, bs.buffer.Len()))
}

func (bs *basicSound) stop() {}

// loopingSound is a sound that gets played over and over again. You can pause & continue
type loopingSound struct {
	buffer  *beep.Buffer
	control *beep.Ctrl
}

func (bs *loopingSound) play(mixer *beep.Mixer) {
	if nil == bs.control {
		speaker.Lock()
		bs.control = &beep.Ctrl{Streamer: beep.Loop(-1, bs.buffer.Streamer(0, bs.buffer.Len()))}
		mixer.Add(bs.control)
		speaker.Unlock()
	}
}

func (bs *loopingSound) stop() {
	if nil != bs.control {
		speaker.Lock()
		bs.control.Streamer = nil
		bs.control = nil
		speaker.Unlock()
	}
}
