package assets

import "github.com/gopxl/beep"

type sound interface {
	decreaseVolume()
	play(mixer *beep.Mixer)
	stop()
}
