package ds

import (
	"math"
	"sync"
	"sync/atomic"
)

type RWMutex struct {
	mutex        *sync.Mutex
	readerWaitCh chan int
	writerWaitCh chan int
	readerCount  int64
	readerWait   int64
}

func (rw *RWMutex) Lock() {
	rw.mutex.Lock()                                                        //加写锁
	if !atomic.CompareAndSwapInt64(&rw.readerCount, 0, -math.MaxInt64-1) { //有读协程,阻塞写协程
		atomic.AddInt64(&rw.readerWait, 1)
		<-rw.writerWaitCh
	}
}

func (rw *RWMutex) Unlock() {
	for i := 0; i < int(atomic.AddInt64(&rw.readerCount, math.MaxInt64)); i++ { //释放信号量，唤醒阻塞的读协程
		rw.readerWaitCh <- 1
	}
	rw.mutex.Unlock()
}

func (rw *RWMutex) RLock() {
	if atomic.AddInt64(&rw.readerCount, 1) < 0 { //有写协获取了锁，读协程阻塞
		<-rw.readerWaitCh
	}
}

func (rw *RWMutex) RUnlock() {
	if atomic.AddInt64(&rw.readerCount, -1) == 0 {
		for i := 0; i < int(rw.readerWait); i++ {
			rw.writerWaitCh <- 1
		}
	}
}


