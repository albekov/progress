package progress

import (
	"fmt"
	"time"
)

const (
	DefaultBarWidth        = 50
	DefaultRefreshInterval = 200 * time.Millisecond
)

type ProgressOptions struct {
	total           int
	barWidth        int
	refreshInterval time.Duration
}

func WithTotal(total int) func(*ProgressOptions) {
	return func(p *ProgressOptions) {
		p.total = total
	}
}

func WithBarWidth(barWidth int) func(*ProgressOptions) {
	return func(p *ProgressOptions) {
		p.barWidth = barWidth
	}
}

func WithRefreshInterval(refreshInterval time.Duration) func(*ProgressOptions) {
	return func(p *ProgressOptions) {
		p.refreshInterval = refreshInterval
	}
}

type Progress struct {
	options          *ProgressOptions
	renderer         Renderer
	current          int
	lastRefresh      time.Time
	lastRefreshValue int
	speed            float64
	started          time.Time
	timeFromStart    time.Duration
}

func New(options ...func(*ProgressOptions)) *Progress {
	progressOptions := &ProgressOptions{
		total:           0,
		barWidth:        DefaultBarWidth,
		refreshInterval: DefaultRefreshInterval,
	}
	for _, option := range options {
		option(progressOptions)
	}
	renderer := NewStdoutRenderer(progressOptions.barWidth)
	progress := &Progress{
		options:  progressOptions,
		renderer: renderer,
		current:  0,
		started:  time.Now(),
	}
	return progress
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
	p.timeFromStart = now.Sub(p.started)
	if p.lastRefresh.IsZero() {
		p.lastRefresh = now
		p.lastRefreshValue = p.current
	} else {
		timeFromLastRefresh := now.Sub(p.lastRefresh)
		if timeFromLastRefresh >= p.options.refreshInterval {
			p.speed = float64(p.current-p.lastRefreshValue) / timeFromLastRefresh.Seconds()
			p.lastRefresh = now
			p.lastRefreshValue = p.current
			p.refresh()
		}
	}
}

func (p *Progress) refresh() {
	p.renderer.Render(p.state())
}

func (p *Progress) state() *ProgressState {
	var percent float64
	var timeRemaining time.Duration
	if p.options.total > 0 {
		percent = float64(p.current) / float64(p.options.total)
		percent = clamp(percent, 0, 1)
		timeRemaining = time.Duration(float64(p.timeFromStart) * (1 - percent) / percent)
	}
	return &ProgressState{
		Current:       p.current,
		Total:         p.options.total,
		Percent:       percent,
		TimeElapsed:   p.timeFromStart,
		TimeRemaining: timeRemaining,
		Speed:         p.speed,
	}
}
