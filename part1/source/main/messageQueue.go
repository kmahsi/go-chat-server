package main

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
