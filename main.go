package main

import (
	"fmt"
	"github.com/aljorhythm/sapere-server/business"
	"github.com/aljorhythm/sapere-server/data/inmemory"
	"log"
	"net/http"
)

func loadDataRetriever() business.BusinessData {
	// dataRetriever := nosql.DataAccess{} real DB connection
	dataRetriever := inmemory.DataAccess{}
	return dataRetriever
}

func main() {
	busi := business.RealBusiness{
		Data: loadDataRetriever(),
	}
	baseHandler := BaseHandler{
		busi,
	}
	http.HandleFunc("/", baseHandler.HelloWorld)
	http.HandleFunc("/feed", baseHandler.Feed)
	http.HandleFunc("/user1", baseHandler.User)
	http.HandleFunc("/ping", baseHandler.Ping)

	port := 8080
	log.Printf("Starting server %d \n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}
