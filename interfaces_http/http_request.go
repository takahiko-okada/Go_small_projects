package main

import "net/http"


func main() {
  resp, err := http.Get("http://google.com")
}
