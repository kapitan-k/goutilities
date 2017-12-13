package concurrency

import (
	"runtime"
	"sync"
)

// Worker represents a single worker.
type Worker struct {
	RunCh chan func()
}

// CreateWorker returns a Worker object wtih runChCap channel size.
func CreateWorker(runChCap int) Worker {
	return Worker{
		RunCh: make(chan func(), runChCap),
	}
}

// Start runs the worker with a new go routine.
func (self Worker) Start(wgDone *sync.WaitGroup, lockOSThread bool) {
	go func() {
		if lockOSThread {
			runtime.LockOSThread()
		}

		for {
			fn := <-self.RunCh
			if fn != nil {
				fn()
			} else {
				return
			}

		}
	}()
	wgDone.Done()
}

// WorkerPool is a pool of workers.
type WorkerPool struct {
	workers []*Worker
	ch      chan func()
	wg      sync.WaitGroup
}

// NewWorkerPool returns a new allocated WorkerPool.
func NewWorkerPool(self *WorkerPool) {
	self = &WorkerPool{}
}

// Init initializes the WorkerPool.
func (self *WorkerPool) Init(workerCnt, capChIn uint64, lockOSThreads bool) (err error) {
	self.workers = make([]*Worker, workerCnt)
	ch := make(chan func(), capChIn)
	self.wg.Add(int(workerCnt))
	for i := 0; i < int(workerCnt); i++ {
		w := &Worker{
			RunCh: ch,
		}
		self.workers[i] = w
		w.Start(&self.wg, lockOSThreads)
	}
	self.ch = ch

	return
}

// Close closes the WorkerPool.
func (self *WorkerPool) Close() {
	close(self.ch)
	self.wg.Wait()
}

// Put puts work to the WorkerPools channel.
func (self *WorkerPool) Put(work func()) {
	self.ch <- work
}

// Offer offers work to the WorkerPools channel.
func (self *WorkerPool) Offer(work func()) bool {
	select {
	case self.ch <- work:
		return true
	default:
	}

	return false
}
