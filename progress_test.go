package progress

import (
	"testing"
	"time"
)

func TestFormatBarTime(t *testing.T) {
	progress := New()
	progress.Set(25)
	progress.timeFromStart = 25 * time.Second
	if progress.formatBarTime() != "00:25<?" {
		t.Errorf("formatBarTime() = %v", progress.formatBarTime())
	}
	progress.Total = 100
	if progress.formatBarTime() != "00:25<01:15" {
		t.Errorf("formatBarTime() = %v", progress.formatBarTime())
	}
}

func TestFormatBar(t *testing.T) {
	progress := New(WithTotal(100), WithBarWidth(8))
	progress.Set(25)
	progress.timeFromStart = 25 * time.Second
	progress.speed = 1
	if progress.renderBar() != "25%  [##      ] 25 / 100 (00:25<01:15, 1.00 it/s)" {
		t.Errorf("formatBar() = %v", progress.renderBar())
	}
}
