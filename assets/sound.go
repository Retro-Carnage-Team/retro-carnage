package assets

import "github.com/faiface/beep"

type sound interface {
	decreaseVolume()
	play(mixer *beep.Mixer)
	stop()
}
