package marathon

import (
	"log"
	"net/http"
	"sync"

	"github.com/SlootSantos/marathon/pkg/client"
)

type Message interface {
	Notify()
}

type messageChannel chan Message

func Listen(ch chan Message) {
	go defaultBroker.listenMsg(ch)
}

var defaultBroker = New(&BrokerNewParams{})

func Handle(w http.ResponseWriter, req *http.Request) {
	defaultBroker.ServeHTTP(w, req)
}

func (b *Broker) listenMsg(ch messageChannel) {
	for {
		msg := <-ch

		msg.Notify()
		b.notifyClients()
	}
}

func (b *Broker) notifyClients() {
	b.subscriptions.mux.Lock()
	defer b.subscriptions.mux.Unlock()

	for _, client := range b.subscriptions.subs {
		if client != nil {
			client.Notification <- "go tigers!"
		}
	}
}

// Broker is the mechanism that handles incoming requests
// and notifies subscriptions when there is new events
type Broker struct {
	Buffer         chan string
	subscriptions  subcriptions
	handlerFunc    http.HandlerFunc
	initalResponse string
}

type subcriptions struct {
	mux  sync.Mutex
	subs []*client.Client
}

type BrokerNewParams struct {
	HandlerFunc    http.HandlerFunc
	InitalResponse string
}

// New creates a new Broker containing an empty list of subscriptions
// and also initates the buffer for events
func New(p *BrokerNewParams) *Broker {
	buf := make(chan string)

	b := &Broker{
		initalResponse: p.InitalResponse,
		handlerFunc:    p.HandlerFunc,
		Buffer:         buf,
		subscriptions: subcriptions{
			mux:  sync.Mutex{},
			subs: []*client.Client{},
		},
	}

	return b
}

func (b *Broker) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Fatal("This ResponseWriter cannot handle Server Sent Events. ReponseWriter needs to implement Flusher Interface")
	}

	if b.handlerFunc != nil {
		b.handlerFunc(w, req)
	}

	newC := client.New(&client.ClientNewParams{
		InitalResponse: b.initalResponse,
		Writer:         setSSEHeaders(w),
		Flusher:        flusher,
	})

	b.subscribe(newC)

	<-req.Context().Done()

	b.unsubscribe(newC)
}

func (b *Broker) subscribe(c *client.Client) {
	b.subscriptions.mux.Lock()
	defer b.subscriptions.mux.Unlock()

	b.subscriptions.subs = append(b.subscriptions.subs, c)
}

func (b *Broker) unsubscribe(c *client.Client) {
	b.subscriptions.mux.Lock()
	defer b.subscriptions.mux.Unlock()

	for idx, c := range b.subscriptions.subs {
		if c == nil || c.ID == c.ID {
			b.subscriptions.subs[idx] = nil
		}
	}
}
