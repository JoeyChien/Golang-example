package main

import (
    "fmt"
    "net/http"
    "log"
    "encoding/json"
)

type Article struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) { 
    articles := Articles{
        Article{
            Title: "My title", Desc: "My desc", Content: "My content"},
    }
    fmt.Fprintf(w, "All articles")
    json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintf(w, "Hello go!") 
}

func handleRequest()  {
    http.HandleFunc("/", homePage) //設定路徑
    http.HandleFunc("/articles", allArticles) //設定路徑
    log.Fatal(http.ListenAndServe(":8080", nil)) //設定port
}

func main() {
    handleRequest()
}