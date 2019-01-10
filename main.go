package main

import (
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	log := log.New(os.Stderr, "[qotm] ", log.LstdFlags)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	quotes := []string{
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

	r := http.NewServeMux()

	r.HandleFunc("/", RichStatusHandler(hostname, quotes))
	r.HandleFunc("/favicon.ico", FileHandler("./favicon.ico"))

	httpServer := &http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	log.Printf("http server listening at %s", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("http server closed unexpectedly: %v\n", err)
	}
}
