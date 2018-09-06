package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request)  {

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Hello from ", name)

}

func main()  {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}