package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/BlackBoyZoovie/store/pkg/routes"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	http.Handle("/", r)
	fmt.Println("starting server at port 8080...")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, r)
	if err != nil {
		log.Fatal("Error starting server at port 8080: ", err)
	}
}