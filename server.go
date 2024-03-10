package inutil

import "github.com/gorilla/mux"

var Server *mux.Router

func StartServer() {
	Server = mux.NewRouter()
}

func Oi() {
	Print("oi")
}
