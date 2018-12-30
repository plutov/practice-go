package collage

import "sync"

// Progress counter
type Progress struct {
	max      int
	cur      int
	callback func(current, max int)

	mu sync.Mutex
}

// NewProgress new progress counter
func NewProgress(max int, callback func(int, int)) *Progress {
	p := &Progress{
		max:      max,
		callback: callback,
	}
	return p
}

// Inc progress
func (p *Progress) Inc() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.cur++

	if p.cur&0xffffff == 0xffffff {
		p.callback(p.cur, p.max)
	}
}
