package syncx

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	count := 0
	p := NewPool[[]byte](func() []byte {
		count += 1
		res := make([]byte, 1, 12)
		res[0] = 'A'
		return res
	})
	res := p.Get()
	assert.Equal(t, "A", string(res))

	res = append(res, 'B')
	p.Put(res)
	res = p.Get()
	if count == 1 {
		assert.Equal(t, "AB", string(res))
	} else {
		assert.Equal(t, "A", string(res))
	}
}

// goos: windows
// goarch: amd64
// pkg: Prove/generalization_tool/syncx
// cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
// BenchmarkPool_Get
// BenchmarkPool_Get/Pool
// BenchmarkPool_Get/Pool-8                12846768                90.29 ns/op
// BenchmarkPool_Get/sync.Pool
// BenchmarkPool_Get/sync.Pool-8          14284489                85.42 ns/op
func BenchmarkPool_Get(b *testing.B) {
	p1 := NewPool[string](func() string {
		return ""
	})

	p2 := &sync.Pool{
		New: func() any {
			return ""
		},
	}
	b.Run("Pool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			p1.Get()
		}
	})
	b.Run("sync.Pool", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			p2.Get()
		}
	})
}
