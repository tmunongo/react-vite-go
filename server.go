package main

import (
  "log"
  "net/http"
)

func main() {
	  fs := http.FileServer(http.Dir("./public"))

    http.Handle("/", fs)
    
    log.Print("listening on localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
