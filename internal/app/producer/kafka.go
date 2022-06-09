package producer

import (
	"context"
	"log"
	"sync"
	"time"

	"gitlab.ozon.dev/betkin/device-api/internal/app/repo"

	"github.com/gammazero/workerpool"
	"gitlab.ozon.dev/betkin/device-api/internal/app/sender"
	"gitlab.ozon.dev/betkin/device-api/internal/model"
)

type Producer interface {
	Start()
	Close()
}

type producer struct {
	n       uint64
	timeout time.Duration

	sender sender.EventSender
	repo   repo.EventRepo
	events <-chan model.DeviceEvent

	workerPool *workerpool.WorkerPool

	wg   *sync.WaitGroup
	done chan bool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	repo repo.EventRepo,
	events <-chan model.DeviceEvent,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}
	done := make(chan bool)

	return &producer{
		n:          n,
		sender:     sender,
		repo:       repo,
		events:     events,
		workerPool: workerPool,
		wg:         wg,
		done:       done,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case event := <-p.events:
					if err := p.sender.Send(&event); err != nil {
						log.Printf("Event %v send failed", event.ID)
						p.workerPool.Submit(func() {
							if err := p.repo.Unlock(context.TODO(), []uint64{event.ID}); err != nil {
								log.Printf("Event %v unlock failed", event.ID)
							}
						})
					} else {
						p.workerPool.Submit(func() {
							if err := p.repo.Remove(context.TODO(), []uint64{event.ID}); err != nil {
								log.Printf("Event %v remove failed", event.ID)
							}
						})
					}
				case <-p.done:
					return
				}
			}
		}()
	}
}

func (p *producer) Close() {
	close(p.done)
	p.wg.Wait()
}
