package main

import (
  "io"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", dog)
  http.HandleFunc("/toby.jpg", dogPic)
  // the first HandleFunc serves dog image, the sescond HandleFunc
  // then opens toby.jpg file
  http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  io.WriteString(w, `<img src="/toby.jpg">`)
}


func dogPic(w http.ResponseWriter, req *http.Request) {
  f, err := os.Open("toby.jpg")
  if err != nil {
    http.Error(w, "file not found", 404)
    return
  }
  defer f.Close()
  // defer ensures that a funcatin is performed later in a
  // program's execution, for purposes of cleanup (ex. closing the file).

  io.Copy(w, f)
}
