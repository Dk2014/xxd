// go test -test.bench=. -test.benchtime=3s

package time

import (
	"fmt"
	"testing"
	"time"
	"sync"
)

func BenchmarkGoTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now().Unix()
	}
	//fmt.Println("time.Now().Unix():", time.Now().Unix(), "VS. AtomicTime:", GetNowTime())
}

func BenchmarkAtomicTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetNowTime()
	}
	//fmt.Println("time.Now().Unix():", time.Now().Unix(), "VS. AtomicTime:", GetNowTime())
}

func BenchmarkConcurrentAtomicTime(b *testing.B) {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			for i := 0; i < b.N; i++ {
				n := GetNowTime()
				if (n > (time.Now().Unix()+1) || n < (time.Now().Unix()-1)) {
					fmt.Println("Error AtomicTime:", n, "time.Now().Unix():", time.Now().Unix())
					b.Fail()
					return
				}
				time.Sleep(1 * time.Millisecond)
			}
		}()
	}

	wg.Wait()
}
