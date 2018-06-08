package calc

import (
	"sync"
	"sync/atomic"
	)

var (
	fibMemorize sync.Map
	fibMax int64
)

func init() {
	fibMemorize.Store(int64(0), int64(0))
	fibMemorize.Store(int64(1), int64(1))

	atomic.StoreInt64(&fibMax, 1)
}

// Calculates a fibonacci number using iteration implementation and memoization approaches.
// Doesn't check overflow, but will return 'err' when implementing it.
func Fib(n uint64) (int64, error) {
	sn := int64(n)

	if v, ok := fibMemorize.Load(sn); ok {
		return v.(int64), nil
	}

	for i := atomic.LoadInt64(&fibMax)+1; i<=sn; i++ {
		if _, ok := fibMemorize.Load(i); ok {
			continue
		}

		v1, _ := fibMemorize.Load(i-1)
		v2, _ := fibMemorize.Load(i-2)
		fibMemorize.Store(i, v1.(int64) + v2.(int64))
	}

	atomic.StoreInt64(&fibMax, sn)

	v, _ := fibMemorize.Load(sn)
	return v.(int64), nil
}
