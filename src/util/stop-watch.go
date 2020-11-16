package util

import (
	"errors"
	"fmt"
	"time"
)

type StopWatch struct {
	startedAt *time.Time
	Duration  int64
	Name      string
}

func (sw *StopWatch) Start() {
	now := time.Now()
	sw.startedAt = &now
}

func (sw *StopWatch) Stop() error {
	if nil == sw.startedAt {
		return errors.New("can't stop a watch that hasn't been started")
	}

	sw.Duration = time.Since(*sw.startedAt).Milliseconds()
	sw.startedAt = nil
	return nil
}

func (sw *StopWatch) DebugMessage() string {
	if sw.Name != "" {
		return fmt.Sprintf("%s took %d ms", sw.Name, sw.Duration)
	}
	return fmt.Sprintf("Task took %d ms", sw.Duration)
}
