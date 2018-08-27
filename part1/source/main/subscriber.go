package main

import (
	"fmt"
	"time"
)

func subscriber12(M MessageQueueInterface) {
	for {
		if x, ok := M.receive(0,0); ok {
			fmt.Println("                                                                          receive:  ", time.Now(), x.message, 1)
		} else {
			break
		}

	}
	time.Sleep(2000 * time.Millisecond)
}

func subscriber11(M MessageQueueInterface) {
	for {
		if x, ok := M.receive(0,1); ok {
			fmt.Println("                                                                          receive:  ", time.Now(), x.message, 2)
		} else {
			break
		}

	}
	time.Sleep(2000 * time.Millisecond)
}

func subscriber21(M MessageQueueInterface) {
	for {
		if x, ok := M.receive(1,0); ok {
			fmt.Println("                                                                          receive:  ", time.Now(), x.message, 3)
		} else {
			break
		}

	}
	time.Sleep(2000 * time.Millisecond)
}

func subscriber22(M MessageQueueInterface) {
	for {
		if x, ok := M.receive(1,1); ok {
			fmt.Println("                                                                          receive:  ", time.Now(), x.message, 4)
		} else {
			break
		}

	}
	time.Sleep(2000 * time.Millisecond)
}
