package main

import (
	"fmt"
	"io"
	"net/http"

	marathon "github.com/SlootSantos/marathon/pkg/broker"
)

var msgChan = make(chan marathon.Message)

func main() {
	marathon.Listen(msgChan)
	http.HandleFunc("/sse2", marathon.Handle)
	http.HandleFunc("/send", handle)

	http.ListenAndServe(":9999", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	msgChan <- &myOwnMessage{
		Val: req.UserAgent(),
	}

	io.WriteString(w, "yiha")
}

type myOwnMessage struct {
	Val string
}

func (m *myOwnMessage) Notify() {
	fmt.Println("Notify!!", m.Val)
}
