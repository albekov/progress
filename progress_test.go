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
	progress.options.total = 100
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

func TestWithRefreshInterval(t *testing.T) {
	progress := New(WithRefreshInterval(500 * time.Millisecond))
	if progress.options.refreshInterval != 500*time.Millisecond {
		t.Errorf("WithRefreshInterval() = %v", progress.options.refreshInterval)
	}
}

func TestSet(t *testing.T) {
	progress := New()
	progress.Set(25)
	if progress.current != 25 {
		t.Errorf("Set() = %v", progress.current)
	}
}

func TestInc(t *testing.T) {
	progress := New()
	progress.Inc()
	progress.Inc()
	if progress.current != 2 {
		t.Errorf("Inc() = %v", progress.current)
	}
}
