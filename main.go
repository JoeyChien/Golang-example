package main

import (
    "fmt"
    "net/http"
    "log"
)

func homePage(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintf(w, "Hello go!") 
}

func handleRequest()  {
    http.HandleFunc("/", homePage) //設定路徑
    log.Fatal(http.ListenAndServe(":8080", nil)) //設定port
}

func main() {
    handleRequest()
}