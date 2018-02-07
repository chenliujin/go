package main

import (
    "fmt"
    "log"
    "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world!")
}

func main() {
    http.HandleFunc("/", sayHello)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe faild")
    }
}
