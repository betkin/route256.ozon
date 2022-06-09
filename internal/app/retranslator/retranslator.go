package retranslator

import (
	"time"

	"gitlab.ozon.dev/betkin/device-api/internal/app/repo"

	"github.com/gammazero/workerpool"
	"gitlab.ozon.dev/betkin/device-api/internal/app/consumer"
	"gitlab.ozon.dev/betkin/device-api/internal/app/producer"
	"gitlab.ozon.dev/betkin/device-api/internal/app/sender"
	"gitlab.ozon.dev/betkin/device-api/internal/model"
)

type Retranslator interface {
	Start()
	Close()
}

type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount uint64
	WorkerCount   int

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type retranslator struct {
	events     chan model.DeviceEvent
	consumer   consumer.Consumer
	producer   producer.Producer
	workerPool *workerpool.WorkerPool
}

func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.DeviceEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	c := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events)
	p := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.Sender,
		cfg.Repo,
		events,
		workerPool)

	return &retranslator{
		events:     events,
		consumer:   c,
		producer:   p,
		workerPool: workerPool,
	}
}

func (r *retranslator) Start() {
	r.producer.Start()
	r.consumer.Start()
}

func (r *retranslator) Close() {
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
}
