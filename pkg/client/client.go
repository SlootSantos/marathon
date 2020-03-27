package client

import (
	"fmt"
	"net/http"
)

// Client is the construct that is responsible for communicating over the HTTP connection
type Client struct {
	ID           string
	writer       http.ResponseWriter
	flusher      http.Flusher
	Notification chan string
}

type ClientNewParams struct {
	Writer         http.ResponseWriter
	Flusher        http.Flusher
	InitalResponse string
}

// New creates a new SSE Client
func New(p *ClientNewParams) *Client {
	clientAckMsg := "ACK SSE"
	if p.InitalResponse != "" {
		clientAckMsg = p.InitalResponse
	}

	client := &Client{ID: RandSeq(10),
		writer:       p.Writer,
		flusher:      p.Flusher,
		Notification: make(chan string),
	}

	client.send(clientAckMsg)
	go client.listen()

	return client
}

func (c *Client) listen() {
	for {
		c.send(<-c.Notification)
	}
}

func (c *Client) send(data string) {
	fmt.Fprintf(c.writer, "data: %s\n\n", data)
	c.flusher.Flush()
}
