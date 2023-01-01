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

	assert.InDelta(t, 250, stopWatch.totalDuration, 50)
}

func TestStopNamedTask(t *testing.T) {
	var stopWatch = &StopWatch{Name: "Buffering sound effect MP5.mp3", totalDuration: 250, startedTimes: 1}
	assert.Equal(t, "1 executions of Buffering sound effect MP5.mp3 took 250 ms (avg)", stopWatch.PrintDebugMessage())
}
