package main

import (
	"fmt"
	"go-tweety/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("tweety-api is running on port 8000")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/signUp", handler.SignUp)
	router.HandleFunc("/showTweets", handler.ShowTweets)
	router.HandleFunc("/newTweet", handler.NewTweet)
	log.Fatal(http.ListenAndServe(":8000", router))
}
