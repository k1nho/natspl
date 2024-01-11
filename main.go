package main

import (
	"github.com/nats-io/nats.go"
)

func main() {

	// connect to nats server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic("could not create nats connection")
	}

	sub := Subscriber{nc: nc}
	pub := Publisher{nc: nc}

	go sub.StartSub()
	pub.StartPublisher()
	pub.SendMessage("another message")
	pub.SendMessage("layer 7")

}
