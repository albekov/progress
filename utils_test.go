package progress

import (
	"testing"
	"time"
)

func TestClamp(t *testing.T) {
	var tests = []struct {
		value, min, max, expected float64
	}{
		{-1, 0, 5, 0},
		{0, 0, 5, 0},
		{1, 0, 5, 1},
		{5, 0, 5, 5},
		{6, 0, 5, 5},
	}
	for _, test := range tests {
		if got := clamp(test.value, test.min, test.max); got != test.expected {
			t.Errorf("clamp(%v, %v, %v) = %v", test.value, test.min, test.max, got)
		}
	}
}

func TestFormatDuration(t *testing.T) {
	var tests = []struct {
		duration time.Duration
		expected string
	}{
		{0, "00:00"},
		{time.Second, "00:01"},
		{time.Minute, "01:00"},
		{time.Hour, "01:00:00"},
		{24 * time.Hour, "1d 00:00:00"},
		{25 * time.Hour, "1d 01:00:00"},
	}
	for _, test := range tests {
		if got := formatDuration(test.duration); got != test.expected {
			t.Errorf("formatDuration(%v) = %v", test.duration, got)
		}
	}
}
