package lb

import "sync/atomic"

// lcBackend 包装 lb.Backend 并记录当前活跃连接数
type lcBackend struct {
	Backend *Backend
	active  int64
}

type LeastConnections struct {
	backends []*lcBackend
}

// NewLeastConnections 构造器，会过滤掉不可用的后端
func NewLeastConnections(backends []*Backend) Picker {
	bs := FilterAlive(backends)
	lcb := make([]*lcBackend, 0, len(bs))
	for _, b := range bs {
		lcb = append(lcb, &lcBackend{Backend: b, active: 0})
	}
	return &LeastConnections{backends: lcb}
}

// Next 返回当前最少活跃连接的后端（不带副作用）
func (l *LeastConnections) Next() *Backend {
	if len(l.backends) == 0 {
		return nil
	}
	// 选择 active 最小的后端，平局时按 slice 顺序
	minIdx := 0
	minVal := atomic.LoadInt64(&l.backends[0].active)
	for i := 1; i < len(l.backends); i++ {
		v := atomic.LoadInt64(&l.backends[i].active)
		if v < minVal {
			minVal = v
			minIdx = i
		}
	}
	return l.backends[minIdx].Backend
}

// Acquire 在 router 确定将要使用某个后端时调用，递增活跃连接计数
func (l *LeastConnections) Acquire(b *Backend) {
	if b == nil {
		return
	}
	for _, lbk := range l.backends {
		if lbk.Backend == b {
			atomic.AddInt64(&lbk.active, 1)
			return
		}
	}
}

// Release 在请求完成时调用，递减活跃连接计数（且不允许小于 0）
func (l *LeastConnections) Release(b *Backend) {
	if b == nil {
		return
	}
	for _, lbk := range l.backends {
		if lbk.Backend == b {
			v := atomic.AddInt64(&lbk.active, -1)
			if v < 0 {
				// 修正为 0，避免计数变负
				atomic.StoreInt64(&lbk.active, 0)
			}
			return
		}
	}
}

// Reset 将所有计数归零（用于测试或维护）
func (l *LeastConnections) Reset() {
	for _, lbk := range l.backends {
		atomic.StoreInt64(&lbk.active, 0)
	}
}

func init() {
	Register("least_connections", NewLeastConnections)
}
