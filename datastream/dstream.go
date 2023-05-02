package datastream

import (
	"fmt"
	"time"
)

type DataStream struct {
	ID   int
	data string
}

func StartStream() {
	stream := make(chan *DataStream)
	go produceStream(stream)
	go consumeStream(stream)
	select {}
}

func produceStream(ds chan *DataStream) {
	for i := 0; ; i++ {
		ds <- &DataStream{
			ID:   i,
			data: fmt.Sprintf("data :%d ", i),
		}
		time.Sleep(time.Second)
	}
}
func consumeStream(ds chan *DataStream) {
	for {
		select {
		case data := <-ds:
			fmt.Println("received data ", data)
		case <-time.After(time.Second * 2):
			fmt.Println("no data received in last 2 seconds")
			return
		}

	}
}
