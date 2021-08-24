package assets

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
)

// loopingSound is a sound that gets played over and over again. You can pause & continue
type loopingSound struct {
	buffer  *beep.Buffer
	control *beep.Ctrl
	volume  *effects.Volume
}

func (bs *loopingSound) play(mixer *beep.Mixer) {
	if nil == bs.control {
		speaker.Lock()
		bs.control = &beep.Ctrl{Streamer: beep.Loop(-1, bs.buffer.Streamer(0, bs.buffer.Len()))}
		bs.volume = &effects.Volume{
			Streamer: bs.control,
			Base:     2,
			Volume:   0,
			Silent:   false,
		}
		mixer.Add(bs.volume)
		speaker.Unlock()
	}
}

func (bs *loopingSound) stop() {
	if nil != bs.control {
		speaker.Lock()
		bs.control.Streamer = nil
		bs.control = nil
		bs.volume = nil
		speaker.Unlock()
	}
}

func (bs *loopingSound) decreaseVolume() {
	speaker.Lock()
	bs.volume.Volume -= volumeChange
	speaker.Unlock()
}
