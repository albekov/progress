package progress

import (
	"testing"
	"time"
)

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
