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

  c := make(chan string)

  for _, link := range links {
    go checkLink(link, c)
    // go spawns new instance of the function in parallel
    // by default one instance, need to change the setting
    // if you want to use multiple thread.
  }

  fmt.Println(<- c)

}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if err != nil {
    fmt.Println(link + "might be down.")
    c <- "Might be down."
    return
  }

  fmt.Println(link, "is up.")
  c <- "is up."

}
