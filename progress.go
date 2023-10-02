package progress

import (
	"fmt"
	"time"
)

type Progress struct {
	Total int

	BarWidth        int
	RefreshInterval time.Duration

	current          int
	lastRefresh      time.Time
	lastRefreshValue int
	speed            float64
}

func New() *Progress {
	return &Progress{
		Total:           0,
		BarWidth:        50,
		RefreshInterval: 200 * time.Millisecond,

		current: 0,
	}
}

func (p *Progress) Set(value int) {
	p.current = value
	p.update()
}

func (p *Progress) Inc() {
	p.Set(p.current + 1)
}

func (p *Progress) Done() {
	p.update()
	p.refresh()
	fmt.Println()
}

func (p *Progress) update() {
	now := time.Now()
	if p.lastRefresh.IsZero() {
		p.lastRefresh = now
		p.lastRefreshValue = p.current
	} else {
		timeFromLastRefresh := now.Sub(p.lastRefresh)
		if timeFromLastRefresh >= p.RefreshInterval {
			p.speed = float64(p.current-p.lastRefreshValue) / timeFromLastRefresh.Seconds()
			p.lastRefresh = now
			p.lastRefreshValue = p.current
			p.refresh()
		}
	}
}

func (p *Progress) refresh() {
	// move cursor to the beginning of the line
	fmt.Print("\r")
	bar := p.renderBar()
	fmt.Print(bar)
	// clear the rest of the line
	fmt.Print("\033[K")
}

func (p *Progress) renderBar() string {
	bar := ""
	if p.Total > 0 {
		percent := float64(p.current) / float64(p.Total)
		percent = clamp(percent, 0, 1)
		bar = fmt.Sprintf("%-4s", fmt.Sprintf("%d%%", int(percent*100)))
		bar += " ["
		barWidth := int(percent * float64(p.BarWidth))
		for i := 0; i < p.BarWidth; i++ {
			if i < barWidth {
				bar += "#"
			} else {
				bar += " "
			}
		}
		bar += "] "
	}
	bar += fmt.Sprintf("%d ", p.current)
	if p.Total > 0 {
		bar += fmt.Sprintf("/ %d ", p.Total)
	}
	bar += fmt.Sprintf("(%.2f it/s)", p.speed)
	return bar
}

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
