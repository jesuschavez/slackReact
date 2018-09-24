package main

import (
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

//FindHandler finds handler function
type FindHandler func(string) (Handler, bool)

//Message struct
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

//Client struct
type Client struct {
	send         chan Message
	socket       *websocket.Conn
	findHandler  FindHandler
	session      *r.Session
	stopChannels map[int]chan bool
}

//NewStopChannel stops go routine
func (c *Client) NewStopChannel(stopKey int) chan bool {
	stop := make(chan bool)
	c.stopChannels[stopKey] = stop
	return stop
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

//func receiver type: Client
func (client *Client) Write() {
	for msg := range client.send {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

//NewClient function
func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	return &Client{
		send:         make(chan Message),
		socket:       socket,
		findHandler:  findHandler,
		session:      session,
		stopChannels: make(map[int]chan bool),
	}
}
