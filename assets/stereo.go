package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"retro-carnage/config"
	"retro-carnage/logging"
	"retro-carnage/util"
	"strings"
	"sync"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

// Stereo is the class we use to play music and sound effects throughout the application.
type Stereo struct {
	effects    map[SoundEffect]sound
	mixer      *beep.Mixer
	music      map[Song]sound
	musicMutex sync.Mutex
	noEffects  bool
	noMusic    bool
}

const (
	mp3Channels   = 2
	mp3Precision  = 2
	mp3SampleRate = 32000
	volumeChange  = 0.3
)

var stereo *Stereo

// NewStereo initializes and returns the Singleton instance of Stereo
func NewStereo() *Stereo {
	if nil == stereo {
		stereo = &Stereo{}
		stereo.initialize()
	}
	return stereo
}

func (sb *Stereo) initialize() {
	var sampleRate beep.SampleRate = mp3SampleRate
	err := speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	if err != nil {
		logging.Error.Println(err.Error())
	}

	sb.mixer = &beep.Mixer{}
	speaker.Play(sb.mixer)

	sb.noEffects = false
	sb.noMusic = false

	var configuration = config.GetConfigService().LoadAudioConfiguration()
	if !configuration.PlayEffects || strings.Contains(os.Getenv("sound"), "no-fx") {
		sb.noEffects = true
	}

	if !configuration.PlayMusic || strings.Contains(os.Getenv("sound"), "no-music") {
		sb.noMusic = true
	}

	sb.effects = make(map[SoundEffect]sound)
	if !sb.noEffects {
		for _, fx := range SoundEffects {
			sound, err := loadSoundEffect(fx)
			if err != nil {
				logging.Error.Panicln(err.Error())
			} else {
				sb.effects[fx] = sound
			}
		}
	}

	sb.music = make(map[Song]sound)
}

// PlayFx starts the playback of a given SoundEffect
func (sb *Stereo) PlayFx(effect SoundEffect) {
	if sb.noEffects {
		return
	}

	var aSound = sb.effects[effect]
	if nil != aSound {
		aSound.play(sb.mixer)
	}
}

// StopFx immediately stops the playback of a given SoundEffect
func (sb *Stereo) StopFx(effect SoundEffect) {
	var aSound = sb.effects[effect]
	if nil != aSound {
		aSound.stop()
	}
}

// IsSongBuffered returns whether or not the given song is buffered (and thus can be played).
func (sb *Stereo) IsSongBuffered(song Song) bool {
	sb.musicMutex.Lock()
	var aSoundIsNil = nil == sb.music[song]
	sb.musicMutex.Unlock()
	return !aSoundIsNil
}

// PlaySong starts the playback of a given Song
func (sb *Stereo) PlaySong(song Song) {
	if sb.noMusic {
		return
	}

	sb.musicMutex.Lock()
	sb.music[song].play(sb.mixer)
	sb.musicMutex.Unlock()
}

// BufferSong loads the given Song but doesn't play it
func (sb *Stereo) BufferSong(song Song) {
	if sb.noMusic {
		return
	}

	if !sb.IsSongBuffered(song) {
		var err error = nil
		var aSound sound
		aSound, err = loadMusic(song)
		if err != nil {
			logging.Error.Panicln(err.Error())
		}

		sb.musicMutex.Lock()
		sb.music[song] = aSound
		sb.musicMutex.Unlock()
	}
}

// StopSong immediately stops the playback of a given Song
func (sb *Stereo) StopSong(song Song) {
	sb.musicMutex.Lock()
	var aSound = sb.music[song]
	if nil != aSound {
		aSound.stop()
	}
	sb.musicMutex.Unlock()
}

// DecreaseVolume decreases the volume of a given Song. Volume will get reset when the Song gets player again.
func (sb *Stereo) DecreaseVolume(song Song) {
	sb.musicMutex.Lock()
	var aSound = sb.music[song]
	if nil != aSound {
		aSound.decreaseVolume()
	}
	sb.musicMutex.Unlock()
}

func loadSoundEffect(fx SoundEffect) (sound, error) {
	stopWatch := util.StopWatch{Name: "Buffering sound effect: " + string(fx)}
	stopWatch.Start()

	var filePath = filepath.Join(".", "sounds", "fx", string(fx))
	buffer, err := readMp3IntoBuffer(filePath)
	if err != nil {
		return nil, err
	}

	_ = stopWatch.Stop()
	logging.Trace.Println(stopWatch.PrintDebugMessage())

	if isLoopingEffect(fx) {
		return &loopingSound{buffer: buffer}, nil
	}
	return &basicSound{buffer: buffer}, nil
}

func loadMusic(song Song) (sound, error) {
	stopWatch := util.StopWatch{Name: "Buffering music: " + string(song)}
	stopWatch.Start()

	var filePath = filepath.Join(".", "sounds", "music", string(song))
	buffer, err := readMp3IntoBuffer(filePath)
	if err != nil {
		return nil, err
	}

	_ = stopWatch.Stop()
	logging.Trace.Println(stopWatch.PrintDebugMessage())

	return &loopingSound{buffer: buffer}, nil
}

func readMp3IntoBuffer(filePath string) (*beep.Buffer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load music from %s: %v", filePath, err)
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode music from %s: %v", filePath, err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	_ = streamer.Close()

	return buffer, nil
}

func isLoopingEffect(fx SoundEffect) bool {
	for _, v := range LoopingSoundEffects {
		if v == fx {
			return true
		}
	}
	return false
}
