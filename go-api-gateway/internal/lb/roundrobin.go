package lb

import (
	"sync/atomic"
)

// RoundRobin 实现（简单循环）
type RoundRobin struct {
	backends []*Backend
	counter  uint64
}

func NewRoundRobin(backends []*Backend) Picker {
	bs := FilterAlive(backends)
	return &RoundRobin{backends: bs}
}

func (r *RoundRobin) Next() *Backend {
	if len(r.backends) == 0 {
		return nil
	}
	// 原子自增，避免竞争
	n := atomic.AddUint64(&r.counter, 1)
	idx := int((n - 1) % uint64(len(r.backends)))
	return r.backends[idx]
}

func init() {
	Register("round_robin", NewRoundRobin)
}
