package main

import (
	"flag"
	"log"
	"net/http"
	numbersApi "number-service/api"
)

/*
http://localhost:8080/numbers?u=http://localhost:8090/rand&u=http://localhost:8090/primes&u=http://localhost:8090/fibo&u=http://localhost:8090/odd
*/

func main() {
	listenAddr := flag.String("http.addr", ":8080", "http listen address")

	flag.Parse()

	http.HandleFunc("/numbers", numbersApi.Get())

	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
