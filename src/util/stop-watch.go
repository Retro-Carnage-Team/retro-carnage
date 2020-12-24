package util

import (
	"errors"
	"fmt"
	"time"
)

type StopWatch struct {
	Name          string
	startedAt     *time.Time
	startedTimes  int
	totalDuration int64
}

func (sw *StopWatch) Start() {
	now := time.Now()
	sw.startedAt = &now
	sw.startedTimes += 1
}

func (sw *StopWatch) Stop() error {
	if nil == sw.startedAt {
		return errors.New("can't stop a watch that hasn't been started")
	}
	sw.totalDuration += time.Since(*sw.startedAt).Milliseconds()
	sw.startedAt = nil
	return nil
}

func (sw *StopWatch) PrintDebugMessage() string {
	if 0 == sw.startedTimes {
		return ""
	}

	if sw.Name != "" {
		return fmt.Sprintf("%d executions of %s took %d ms (avg)", sw.startedTimes, sw.Name, sw.totalDuration/int64(sw.startedTimes))
	}
	return fmt.Sprintf("%d executions took %d ms (avg)", sw.startedTimes, sw.totalDuration/int64(sw.startedTimes))
}
