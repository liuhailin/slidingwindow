package slidingwindow

import "sync/atomic"

type LimitProvider struct {
	dynamicLimit int64
}

func (dl *LimitProvider)FetchDynamicLimit() *int64 {
	return &dl.dynamicLimit
	
}

func (dl *LimitProvider)SetDynamicLimit(l int64) {
	atomic.AddInt64(dl.FetchDynamicLimit(),l)
}