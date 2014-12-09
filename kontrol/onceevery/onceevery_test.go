package onceevery

import (
	"sync"
	"testing"
	"time"
)

func TestOnceEvery(t *testing.T) {
	once := New(time.Second)
	count := 0
	var countMu sync.Mutex

	go func() {
		for i := 0; i < 100; i++ {
			once.Do(func() {
				countMu.Lock()
				count++
				countMu.Unlock()
			})
		}
	}()

	time.Sleep(time.Second * 2)

	countMu.Lock()
	defer countMu.Unlock()

	if count != 2 {
		t.Errorf("function should be calle two times, got '%d'", count)
	}

}
