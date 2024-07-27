package GoBasics

import (
	"errors"
	"time"
)

type StopWatch struct {
	startTime time.Time
	endTime   time.Time
}

func (s *StopWatch) Reset() {
	s.startTime = time.Time{}
	s.endTime = time.Time{}
}

func (s *StopWatch) Start() {
	s.startTime = time.Now()
	s.endTime = time.Time{}
	Label("Start: " + s.startTime.String())
}

func (s *StopWatch) Stop() {
	s.endTime = time.Now()
	Label("Stop: " + s.startTime.String())
}

func (s *StopWatch) Elapse() (time.Duration, error) {
	if s.startTime.IsZero() {
		return 0, errors.New("StopWatch never started")
	}
	if s.endTime.IsZero() {
		return 0, errors.New("StopWatch never stopped")
	}
	if s.startTime.After(s.endTime) {
		return 0, errors.New("StopWatch start time is after stop time")
	}

	delta := s.endTime.Sub(s.startTime)
	Label("Delta: " + delta.String())
	return delta, nil
}
