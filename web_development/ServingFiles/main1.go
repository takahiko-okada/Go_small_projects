package main

import (
  "io"
  "net/http"
)

func main() {
  http.HandleFunc("/", dog)
  // func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
  http.ListenAndServe(":8080", nil)
  // func ListenAndServe(addr string, handler Handler) error

func dog(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  io.WriteString(w, `<! -- not serving from local sesrver -->
    <img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}
