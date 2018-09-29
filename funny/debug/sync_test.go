package debug
import (
	sync "github.com/funny/debug/sync"
	"testing"
)


// go test  -bench=. -run BenchmarkMutex   -count=5
// go test -tags deadlock -bench=. -run BenchmarkMutex  -count=5
func BenchmarkMutex(b *testing.B) {
	var m sync.Mutex
	var count uint64
	for i := 0; i < b.N; i++ {
		m.Lock()
		count++
		m.Unlock()
	}
}