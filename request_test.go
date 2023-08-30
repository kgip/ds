package ds

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func(index int) {
			t.Log("start request ---- > " + fmt.Sprintf("%d", index))
			_, err := http.Get("http://localhost:6060/api/user/list")
			time.Sleep(10 * time.Millisecond)
			t.Log(err)
			t.Log("request end")
			wg.Done()
		}(i + 1)
	}
	wg.Wait()
}

func TestGroupRun(t *testing.T) {
	for i := 0; i < 10; i++ {
		index := i
		t.Run(fmt.Sprintf("group%d", i), func(t *testing.T) {
			t.Log(index)
			time.Sleep(time.Second)
		})
	}
}
