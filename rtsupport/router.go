package main

import (
	"fmt"
	"net/http"

	r "github.com/dancannon/gorethink"

	"github.com/gorilla/websocket"
)

//Handler func type
type Handler func(*Client, interface{})

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//NewRouter initializes router object
//like a constructor in js
func NewRouter(session *r.Session) *Router {
	return &Router{
		rules:   make(map[string]Handler),
		session: session,
	}
}

// Router type struct
type Router struct {
	//map: key -> event name string ex: "add channel"  value -> handler function: addChannel()
	rules   map[string]Handler
	session *r.Session
}

// Handle http request
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

//FindHandler returns handler function based on msgName param
func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}

//ServeHTTP interfaces with client
func (e *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//upgrades HTTP request conn to Websocket request conn
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(socket, e.FindHandler, e.session)

	go client.Write()
	client.Read()
}
