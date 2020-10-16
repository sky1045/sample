package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("hell")
	})

	http.HandleFunc("/send", func(w http.ResponseWriter, req *http.Request) {
        resp, err := http.Get("receiver.receiver.svc.cluster.local")
		defer resp.Body.Close()
		log.Println("send")
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Printf("%s\n", string(data))
	})

	log.Println("listening at port 80")
	log.Println(http.ListenAndServe(":80", nil))
}
