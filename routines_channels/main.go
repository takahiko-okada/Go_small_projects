package main

import (
        "fmt"
        "net/http"
        )


func main() {
  links := []string {
    "http://google.com",
    "http://stackoverflow.com",
    "http://facebook.com",
    "http://golang.org",
    "http://amazon.com",
  }

  for _, link := range links {
    go checkLink(link)
    // go spawns new instance of the function in parallel
    // by default one instance, need to change the setting
    // if you want to use multiple thread.
  }
}

func checkLink(link string) {
  _, err := http.Get(link)
  if err != nil {
    fmt.Println(link + "might be down.")
    return
  }

  fmt.Println(link, "is up.")
}
