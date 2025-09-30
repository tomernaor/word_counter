package worker_pool

import (
	"context"
	"sync"
)

type Job func(ctx context.Context, poolId int) error

type Pool struct {
	ch     chan Job
	wg     sync.WaitGroup
	mu     sync.Mutex
	errors []error
}

func (p *Pool) Errors() []error {
	return p.errors
}

func NewPool(ctx context.Context, size int) *Pool {
	p := Pool{
		ch: make(chan Job),
		wg: sync.WaitGroup{},
	}

	for i := 0; i < size; i++ {
		go p.worker(ctx, i, &p.wg, p.ch)
	}

	return &p
}

func (p *Pool) worker(ctx context.Context, id int, wg *sync.WaitGroup, jobs <-chan Job) {
	for job := range jobs {
		if err := job(ctx, id); err != nil {
			p.mu.Lock()
			p.errors = append(p.errors, err)
			p.mu.Unlock()
		}

		wg.Done()
	}
}

func (p *Pool) RunJob(job Job) {
	p.wg.Add(1)
	p.ch <- job
}

func (p *Pool) Close() {
	close(p.ch)
	p.wg.Wait()
}
