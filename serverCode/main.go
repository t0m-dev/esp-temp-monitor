package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func printRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", printRequest)

	log.Fatal(http.ListenAndServe(":4444", mux))

}
