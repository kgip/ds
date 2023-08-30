package ds

import (
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		t.Log(rand.Float64())
		time.Sleep(time.Microsecond)
	}
}
