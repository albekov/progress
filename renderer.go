package progress

import (
	"fmt"
	"time"
)

type ProgressState struct {
	Current       int
	Total         int
	Percent       float64
	TimeElapsed   time.Duration
	TimeRemaining time.Duration
	Speed         float64
}

type Renderer interface {
	Render(state *ProgressState)
}

type StdoutRenderer struct {
	barWidth int
}

func NewStdoutRenderer(barWidth int) *StdoutRenderer {
	return &StdoutRenderer{
		barWidth: barWidth,
	}
}

func (r *StdoutRenderer) Render(state *ProgressState) {
	// move cursor to the beginning of the line
	fmt.Print("\r")
	bar := r.formatProgress(state)
	fmt.Print(bar)
	// clear the rest of the line
	fmt.Print("\033[K")
}

func (r *StdoutRenderer) formatProgress(state *ProgressState) string {
	str := ""
	if state.Total > 0 {
		str = fmt.Sprintf("%-4s", fmt.Sprintf("%d%%", int(state.Percent*100)))
		str += " " + r.formatBar(state) + " "
	}
	str += fmt.Sprintf("%d ", state.Current)
	if state.Total > 0 {
		str += fmt.Sprintf("/ %d ", state.Total)
	}
	str += fmt.Sprintf("(%s, %.2f it/s)", formatProgressTime(state), state.Speed)
	return str
}

func (r *StdoutRenderer) formatBar(state *ProgressState) string {
	str := "["
	completed := int(state.Percent * float64(r.barWidth))
	for i := 0; i < r.barWidth; i++ {
		if i < completed {
			str += "#"
		} else {
			str += " "
		}
	}
	str += "]"
	return str
}

func formatProgressTime(state *ProgressState) string {
	formatted := formatDuration(state.TimeElapsed)
	if state.Total > 0 {
		formatted += fmt.Sprintf("<%s", formatDuration(state.TimeRemaining))
	} else {
		formatted += "<?"
	}
	return formatted
}
