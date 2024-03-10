package inutil

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/pat"
)

var Router *pat.Router
var Server *http.Server

func StartServer(address string) {
	Router = pat.New()

	Server = &http.Server{
		Addr: address, //"0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      Router, // Pass our instance of gorilla/mux in.
	}
}

func Oi() {
	Print("oi2")
}

func RunServer() {
	log.Fatal(Server.ListenAndServe())
}
