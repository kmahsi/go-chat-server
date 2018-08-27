package main

import (
	"fmt"
	"time"
)

func subscriber1(M MessageQueueInterface) {
	for {
		if x, ok := M.receive(0); ok {
			fmt.Println("                                                                          receive:  ", time.Now(), x.message, 1)
		} else {
			break
		}

	}
	time.Sleep(2000 * time.Millisecond)
}

func subscriber2(M MessageQueueInterface) {
	for {
		if x, ok := M.receive(1); ok {
			fmt.Println("                                                                          receive:  ", time.Now(), x.message, 2)
		} else {
			break
		}

	}
	time.Sleep(2000 * time.Millisecond)
}
