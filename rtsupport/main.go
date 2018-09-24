package main

import (
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
)

//Channel struct
type Channel struct {
	Id   string `json:"id" gorethink: "id, omitempty"`
	Name string `json:"name" gorethink: "name"`
}

//User struct

type User struct {
	Id   string `gorethink: "id, omitempty"`
	Name string `gorethink: "name"`
}

func main() {
	//connect to server
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "rtsupport",
	})

	if err != nil {
		log.Panic(err.Error())
	}
	//initalizes new Router
	router := NewRouter(session)

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
