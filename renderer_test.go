package progress

import (
	"testing"
	"time"
)

func TestFormatBarTime(t *testing.T) {
	state := &ProgressState{
		Current:     25,
		TimeElapsed: 25 * time.Second,
	}
	if formatProgressTime(state) != "00:25<?" {
		t.Errorf("formatBarTime() = %v", formatProgressTime(state))
	}
	state.Total = 100
	state.TimeRemaining = 75 * time.Second
	if formatProgressTime(state) != "00:25<01:15" {
		t.Errorf("formatBarTime() = %v", formatProgressTime(state))
	}
}

func TestFormatBar(t *testing.T) {
	renderer := NewStdoutRenderer(8)
	state := &ProgressState{
		Current:       25,
		Total:         100,
		Percent:       0.25,
		TimeElapsed:   25 * time.Second,
		TimeRemaining: 75 * time.Second,
		Speed:         1,
	}
	if renderer.formatProgress(state) != "25%  [##      ] 25 / 100 (00:25<01:15, 1.00 it/s)" {
		t.Errorf("formatBar() = %v", renderer.formatProgress(state))
	}
}
