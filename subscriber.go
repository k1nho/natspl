package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	nc *nats.Conn
}

func (s Subscriber) StartSub() {
	ch := make(chan *nats.Msg, 64)
	sub, err := s.nc.ChanSubscribe(foo, ch)
	if err != nil {
		panic("could not create the channel subscriber")
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case msg := <-ch:
			fmt.Println(string(msg.Data))
		case <-signalCh:
			if err := sub.Unsubscribe(); err != nil {
				panic("oh no we could not unsubscribe")
			}
			close(ch)
			os.Exit(0)
		}
	}

}
