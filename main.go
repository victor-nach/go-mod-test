package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

type Response struct {
	Status int
	Message string
}

func main() {
	router := mux.NewRouter()

	PORT := "3000"
	
	router.HandleFunc("/", welcome)

	log.Println("Serving on port: ", PORT)
	log.Fatal(http.ListenAndServe(":" + PORT, router))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status: 200,
		Message: "Hello World",
	}
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}