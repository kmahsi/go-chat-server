package main

import (
	"time"
)

const NUMBER_OF_CHANNELS int = 7

type MessageStruct struct {
	message   string
	publisher int
}

type MessageQueue struct {
	streamChannel []chan MessageStruct
}


func CreateMessageQueue() *MessageQueue {
	return new(MessageQueue)
}

type MessageQueueInterface interface {
	send(MessageStruct, int)
	receive(int) (MessageStruct, bool)
	addChannel()
}

func (M *MessageQueue) send(s MessageStruct, QueueNum int) {
	M.streamChannel[QueueNum] <- s
}

func (M *MessageQueue) receive(QueueNum int) (MessageStruct, bool) {
	x, ok := <-M.streamChannel[QueueNum]
	return x, ok
}

func (M *MessageQueue) addChannel() {
	M.streamChannel = append(M.streamChannel, make(chan MessageStruct, 100))
}

func main() {
	var messanger MessageQueueInterface

	messanger = CreateMessageQueue()
	for i := 0; i < NUMBER_OF_CHANNELS; i++ {
		messanger.addChannel()
	}
	go publisher1(messanger)
	go publisher2(messanger)
	go publisher3(messanger)
	go subscriber1(messanger)
	go subscriber2(messanger)
	time.Sleep(7000000 * time.Microsecond)
}
