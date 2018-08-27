package main

import (
	"fmt"
	"strconv"
	"time"
)

func publisher1(M MessageQueueInterface) {
	for {
		for i := 0; i < 10; i++ {
			sendMessage(MessageStruct{message: strconv.Itoa(i + 1), publisher: 1}, M, 1000, 0)
		}
	}
}

func publisher2(M MessageQueueInterface) {
	for {
		for i := 0; i < 10; i++ {
			sendMessage(MessageStruct{message: "hello", publisher: 2}, M, 500, 0)
		}
	}
}

func publisher3(M MessageQueueInterface) {
	for {
		for i := 20; i > 0; i-- {
			sendMessage(MessageStruct{message: strconv.Itoa(i), publisher: 3}, M, 700, 1)
		}
	}
}

func sendMessage(message MessageStruct, que MessageQueueInterface, delay time.Duration, QueueNum int) {
	que.send(message, QueueNum)
	fmt.Println("send:    ", time.Now(), message.message, QueueNum)
	time.Sleep(delay * time.Millisecond)
}
