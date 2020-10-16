package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("receive")
        w.Write([]byte("received well"))
	})

	log.Println("listening at port 80")
	log.Println(http.ListenAndServe(":80", nil))
}
