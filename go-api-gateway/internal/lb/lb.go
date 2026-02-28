package lb

import (
	"errors"
	"fmt"
	"sync"
)

// Backend 表示上游后端节点
type Backend struct {
	URL    string
	Weight int
	Alive  bool
}

// Picker 负载均衡器接口，返回被选中的后端（可能为 nil）
type Picker interface {
	Next() *Backend
}

// Factory 用于创建 Picker 的工厂函数
type Factory func(backends []*Backend) Picker

var (
	regMu    sync.RWMutex
	registry = make(map[string]Factory)
)

// Register 向注册表注册一个策略实现
func Register(name string, f Factory) {
	regMu.Lock()
	defer regMu.Unlock()
	registry[name] = f
}

// NewPickerByName 用名字创建 Picker
func NewPickerByName(name string, backends []*Backend) (Picker, error) {
	regMu.RLock()
	f, ok := registry[name]
	regMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("lb strategy not registered: %s", name)
	}
	if len(backends) == 0 {
		return nil, errors.New("no backends provided")
	}
	return f(backends), nil
}

// FilterAlive 返回只包含 Alive 的后端（拷贝）
func FilterAlive(backends []*Backend) []*Backend {
	out := make([]*Backend, 0, len(backends))
	for _, b := range backends {
		if b != nil && b.Alive {
			out = append(out, b)
		}
	}
	return out
}
