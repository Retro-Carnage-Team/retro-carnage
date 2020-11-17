package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStopSomeTask(t *testing.T) {
	var stopWatch = &StopWatch{}
	stopWatch.Start()
	time.Sleep(250 * time.Millisecond)
	stopWatch.Stop()

	assert.InDelta(t, 250, stopWatch.Duration, 50)
}

func TestStopNamedTask(t *testing.T) {
	var stopWatch = &StopWatch{Name: "Buffering sound effect MP5.mp3", Duration: 250}
	assert.Equal(t, "Buffering sound effect MP5.mp3 took 250 ms", stopWatch.DebugMessage())
}
