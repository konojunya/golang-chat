package main

import (
	"flag"
	"log"
	"net/http"
)

// web application address
var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {

	// access log
	log.Println(r.URL)

	// URLが/でなければNot Found
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	// index.htmlを返す
	http.ServeFile(w, r, "index.html")
}

func main() {
	flag.Parse()

	// newHubは構造体のHubのアドレスを返す
	hub := newHub()
	go hub.run()

	// http.HandleFuncでルーティングと関数をハンドリング
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	// *addr = :8080
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
