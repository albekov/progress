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
	started          time.Time
}

func New(options ...func(*Progress)) *Progress {
	progress := &Progress{
		Total:           0,
		BarWidth:        50,
		RefreshInterval: 200 * time.Millisecond,
		current:         0,
		started:         time.Now(),
	}
	for _, option := range options {
		option(progress)
	}
	return progress
}

func WithTotal(total int) func(*Progress) {
	return func(p *Progress) {
		p.Total = total
	}
}

func WithBarWidth(barWidth int) func(*Progress) {
	return func(p *Progress) {
		p.BarWidth = barWidth
	}
}

func WithRefreshInterval(refreshInterval time.Duration) func(*Progress) {
	return func(p *Progress) {
		p.RefreshInterval = refreshInterval
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
	bar += fmt.Sprintf("(%s, %.2f it/s)", p.formatBarTime(), p.speed)
	return bar
}

func (p *Progress) formatBarTime() string {
	timeFromStart := time.Since(p.started)
	formatted := formatDuration(timeFromStart)
	if p.Total > 0 {
		percent := float64(p.current) / float64(p.Total)
		percent = clamp(percent, 0, 1)
		totalTime := time.Duration(float64(timeFromStart) / percent)
		timeLeft := totalTime - timeFromStart
		formatted += fmt.Sprintf("<%s", formatDuration(timeLeft))
	} else {
		formatted += "<?"
	}
	return formatted
}