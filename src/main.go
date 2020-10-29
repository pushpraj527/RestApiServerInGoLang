package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func clientCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Client Creation Started"))
	fmt.Println("Client creation requsted")

	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var clientBody string

	json.Unmarshal(bodyByte, &clientBody)

	fmt.Println("body from client", clientBody)
}

func handleRequests() {
	handlerRouter := mux.NewRouter().StrictSlash(true)
	handlerRouter.HandleFunc("/", clientCreate).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", handlerRouter))
}

func main() {
	handleRequests()
}
