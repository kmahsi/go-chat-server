package main

import (
	"time"
)

const NUMBER_OF_QUEUES int = 2

type MessageStruct struct {
	message   string
	publisher int
}

type MessageQueue struct {
	streamChannels [][]chan MessageStruct
}

func CreateMessageQueue() *MessageQueue {
	return new(MessageQueue)
}

type MessageQueueInterface interface {
	send(MessageStruct, int)
	receive(int, int) (MessageStruct, bool)
	addChannel()
	addSubscriber(int)
}

func (M *MessageQueue) addSubscriber(channelNumber int) {
	M.streamChannels[channelNumber] = append(M.streamChannels[channelNumber], make(chan MessageStruct, 100))
}

func (M *MessageQueue) send(s MessageStruct, QueueNum int) {
	for index := range (M.streamChannels[QueueNum]) {
		M.streamChannels[QueueNum][index] <- s
	}
}

func (M *MessageQueue) receive(QueueNum int, subscriberNum int) (MessageStruct, bool) {
	return MessageStruct{}, false
	select {
	case x := <-M.streamChannels[QueueNum][subscriberNum]:
		return x, true
	default:
		return MessageStruct{}, false
	}
}

func (M *MessageQueue) addChannel() {
	M.streamChannels = append(M.streamChannels, make([]chan MessageStruct, 0))
}

func main() {
	var messanger MessageQueueInterface

	messanger = CreateMessageQueue()
	for i := 0; i < NUMBER_OF_QUEUES; i++ {
		messanger.addChannel()
	}
	messanger.addSubscriber(0)
	messanger.addSubscriber(0)
	messanger.addSubscriber(1)
	messanger.addSubscriber(1)

	go publisher1(messanger)
	go publisher2(messanger)
	go publisher3(messanger)
	go subscriber11(messanger)
	go subscriber12(messanger)
	go subscriber21(messanger)
	go subscriber22(messanger)
	time.Sleep(7000000 * time.Microsecond)
}
