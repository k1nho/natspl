package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	nc *nats.Conn
}

func (p Publisher) StartPublisher() {
	var err error
	for i := 0; i < 5; i++ {
		err = p.nc.Publish(foo, []byte(fmt.Sprintf("message %d", i)))
		if err != nil {
			panic("could not publish to foo")
		}
		time.Sleep(1 * time.Second)
	}
}

func (p Publisher) SendMessage(msg string) {
	err := p.nc.Publish(foo, []byte(msg))
	if err != nil {
		panic("could not publish message")
	}
	time.Sleep(1 * time.Second)
}
