package lb

import "sync"

// 使用 Smooth Weighted Round Robin（Nginx 演算法）
type wBackend struct {
	b       *Backend
	weight  int // 权重（固定值）
	current int // 当前累积值（动态变化）
}

type WeightedRoundRobin struct {
	mu       sync.Mutex
	backends []*wBackend
	total    int // 所有 backend 的权重总和
}

func NewWeightedRoundRobin(backends []*Backend) Picker {
	bs := FilterAlive(backends)
	wbs := make([]*wBackend, 0, len(bs))
	total := 0
	for _, b := range bs {
		w := 1
		if b.Weight > 0 {
			w = b.Weight
		}
		wbs = append(wbs, &wBackend{b: b, weight: w, current: 0})
		total += w
	}
	return &WeightedRoundRobin{backends: wbs, total: total}
}

func (w *WeightedRoundRobin) Next() *Backend {
	w.mu.Lock()
	defer w.mu.Unlock()
	if len(w.backends) == 0 {
		return nil
	}
	// Smooth Weighted Round-Robin
	var best *wBackend
	for _, wb := range w.backends {
		wb.current += wb.weight
		if best == nil || wb.current > best.current {
			best = wb
		}
	}
	if best == nil {
		return nil
	}
	best.current -= w.total
	return best.b
}

func init() {
	Register("weighted_round_robin", NewWeightedRoundRobin)
}
