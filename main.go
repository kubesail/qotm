package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var hostname string
var quotes = []string{
	"Abstraction is ever present.",
	"A late night does not make any sense.",
	"A principal idea is omnipresent, much like candy.",
	"Nihilism gambles with lives, happiness, and even destiny itself!",
	"The light at the end of the tunnel is interdependent on the relatedness of motivation, subcultures, and management.",
	"Utter nonsense is a storyteller without equal.",
	"Non-locality is the driver of truth. By summoning, we vibrate.",
	"A small mercy is nothing at all?",
	"The last sentence you read is often sensible nonsense.",
	"668: The Neighbor of the Beast.",
}

func main() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", GetTime)
	log.Println("QOTM starting")
	http.ListenAndServe(":8000", r)
}

func GetTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{
    "hostname": "%s",
    "ok": true,
    "quote": "%s",
    "time": "%s",
    "version": "2.0"
}`, hostname, quotes[rand.Intn(len(quotes))], time.Now().Format(time.RFC3339))
}
