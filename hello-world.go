package main

import (
        "net/http"
        "log"
)

func main() {
        http.HandleFunc("/", hello)
        log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r  *http.Request) {
        w.Write([]byte("Hello, World!!"))
}
