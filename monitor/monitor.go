package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
	"auto-wallet/monitor/transactions"
	"github.com/pebbe/zmq4"
	"encoding/json"
	"io/ioutil"
)

func Shutdown(timeout time.Duration) {
	select {
	case <-time.After(timeout):
	}
}

type Config struct {
    Addr string
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// print out current zmq version
	major, minor, patch := zmq4.Version()
	fmt.Printf("running ZMQ %d.%d.%d\n", major, minor, patch)

	// read config file
	value := Config{}
	fileBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(fileBytes, &value); err != nil {
		panic(err)
	}
	fmt.Printf("ZMQ addr: %s\n", value.Addr)

	// start feeds
	go transactions.StartTxFeed(value.Addr)
	go transactions.StartMilestoneFeed(value.Addr)
	go transactions.StartConfirmationFeed(value.Addr)

	select {
	case <-sigs:
		Shutdown(time.Duration(1500) * time.Millisecond)
	}
}