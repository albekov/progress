package progress

import (
	"fmt"
	"time"
)

type number interface {
	float64
}

func clamp[T number](value T, min T, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func formatDuration(duration time.Duration) string {
	days := duration / (24 * time.Hour)
	duration = duration % (24 * time.Hour)
	hours := duration / time.Hour
	duration = duration % time.Hour
	minutes := duration / time.Minute
	duration = duration % time.Minute
	seconds := duration / time.Second

	formatted := ""

	if days > 0 {
		formatted += fmt.Sprintf("%dd ", days)
	}

	if hours > 0 {
		formatted += fmt.Sprintf("%02d:", hours)
	}

	formatted += fmt.Sprintf("%02d:%02d", minutes, seconds)

	return formatted
}
