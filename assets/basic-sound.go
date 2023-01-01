package assets

import "github.com/faiface/beep"

// basicSound is a sound that gets played exactly one time. No further interaction is possible
type basicSound struct {
	buffer *beep.Buffer
}

func (bs *basicSound) play(mixer *beep.Mixer) {
	mixer.Add(bs.buffer.Streamer(0, bs.buffer.Len()))
}

func (bs *basicSound) stop() {}

func (bs *basicSound) decreaseVolume() {}
