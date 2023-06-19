package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

func main() {
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "127.0.0.1:10041", Path: "/ws/receive"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("received message: %s", message)
		}
	}()

	for {
		select {
		case <-done:
		case <-time.After(time.Second):
			// err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "123"))
			// if err != nil {
			// 	log.Println("write close:", err)
			// 	return
			// }
			err := c.WriteMessage(1, []byte("abc"))
			if err != nil {
				log.Println(err)
			}
		}
	}
}
