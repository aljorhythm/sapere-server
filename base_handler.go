package main

import (
	"encoding/json"
	"fmt"
	"github.com/aljorhythm/sapere-server/business"
	"log"
	"net/http"
)

type Business interface {
	GetReadingFeed() (business.Feed, error)
	GetUser(i int64) (business.User, error)
}

type BaseHandler struct {
	Business
}

func (h BaseHandler) Feed(w http.ResponseWriter, r *http.Request) {
	feed, error := h.Business.GetReadingFeed()

	if error != nil {
		error := json.NewEncoder(w).Encode(feed)

		if error != nil {
			panic(error)
		}
		return
	}

	err := json.NewEncoder(w).Encode(feed)

	alertIfError(err)
}

func (h BaseHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, you are visiting %s!", r.URL.Path[1:])
}

func (h BaseHandler) User(w http.ResponseWriter, r *http.Request) {
	user, error := h.GetUser(1)
	if error != nil {
		err := json.NewEncoder(w).Encode(user)

		if err != nil {
			panic(err)
		}
		return
	}

	err := json.NewEncoder(w).Encode(user)

	alertIfError(err)
}

func (h BaseHandler) Ping(w http.ResponseWriter, r *http.Request) {
	pingResponse := struct {
		Up bool `json:"up"`
	}{
		Up: true,
	}

	err := json.NewEncoder(w).Encode(pingResponse)

	alertIfError(err)
}

func alertIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
