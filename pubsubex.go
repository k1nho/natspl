package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	err := nc.Publish("greet.joe", []byte("hello"))
	checkErr(err)

	sub, _ := nc.SubscribeSync("greet.*")

	msg, _ := sub.NextMsg(10 * time.Millisecond)
	fmt.Println("subscribe after a publish")
	fmt.Printf("message is nil? %v\n", msg == nil)

	err = nc.Publish("greet.joe", []byte("hello joe"))
	checkErr(err)
	err = nc.Publish("greet.pam", []byte("hello pam"))
	checkErr(err)

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q, msg subject: %q\n", string(msg.Data), msg.Subject)
	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q, msg subject: %q\n", string(msg.Data), msg.Subject)

	err = nc.Publish("greet.bob", []byte("hello bob"))
	checkErr(err)

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q, msg subject: %q\n", string(msg.Data), msg.Subject)

}
