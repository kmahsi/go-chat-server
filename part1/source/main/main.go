package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

func chat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err!= nil{
			log.Println("read:", nil)
			break
		}
		fmt.Println(r)
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil{
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	//var messanger MessageQueueInterface
	//
	//messanger = CreateMessageQueue()
	//for i := 0; i < NUMBER_OF_QUEUES; i++ {
	//	messanger.addChannel()
	//}
	//messanger.addSubscriber(0)
	//messanger.addSubscriber(0)
	//messanger.addSubscriber(1)
	//messanger.addSubscriber(1)

	//go publisher1(messanger)
	//go publisher2(messanger)
	//go publisher3(messanger)
	//go subscriber11(messanger)
	//go subscriber12(messanger)
	//go subscriber21(messanger)
	//go subscriber22(messanger)
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/chat", chat)
	log.Fatal(http.ListenAndServe(*addr, nil))

	time.Sleep(7000000 * time.Microsecond)
}
