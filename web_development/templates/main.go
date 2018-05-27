package main

import (
  "log"
  "os"
  "text/template"
)

// variable tpl can be considered as a container
// of *template(s) which are the memory addresses of the parsed template
var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("tpl/*"))
}

func main() {
  // Execute template, first one?
  err := tpl.Execute(os.Stdout, nil)
  if err != nil {
    log.Fatalln(err)
  }

  // Execute specific template with ExecuteTemplate func
  err = tpl.ExecuteTemplate(os.Stdout, "Test1.gohtml", nil)
  if err != nil {
    log.Fatalln(err)
  }
  // Same as above
  err = tpl.ExecuteTemplate(os.Stdout, "Test2.gohtml", nil)
  if err != nil {
    log.Fatalln(err)
  }
}

// ExecuteTemplate(writer, file, data to pass)
