package ratelimit

import "net/http"

type Limiter interface {
	Allow(r *http.Request) (bool, error)
	Name() string
}
