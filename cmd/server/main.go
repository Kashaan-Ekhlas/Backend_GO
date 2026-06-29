package main

import (
	"log"
	"net/http"
	"github.com/Kashaan-Ekhlas/Backend_GO/internal"
)

func main(){
  mux := internal.NewRouter()

	server:= &http.Server{
		Addr: ":5005",
		Handler: mux,
	}

	log.Println("Listening on port 5005")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
