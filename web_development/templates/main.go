package main

import (
  "log"
  "os"
  "text/template"
)

// variable tpl can be considered as a container
// of *template(s) which are the memory addresses of the parsed template
var tpl *template.Template

type property struct {
  Capital string
  Weather string
}

func init() {
  tpl = template.Must(template.ParseGlob("./web_development/templates/tpl/*"))
}

func main() {
  // Execute template, first one?
  err := tpl.Execute(os.Stdout, "is to find it.")
  if err != nil {
    log.Fatalln(err)
  }

  err = tpl.ExecuteTemplate(os.Stdout, "test0.gohtml", "is to find it.")
  if err != nil {
    log.Fatalln(err)
  }

  // Execute specific template with ExecuteTemplate func
  err = tpl.ExecuteTemplate(os.Stdout, "test1.gohtml", nil)
  if err != nil {
    log.Fatalln(err)
  }
  // Same as above
  err = tpl.ExecuteTemplate(os.Stdout, "test2.gohtml", nil)
  if err != nil {
    log.Fatalln(err)
  }


  // passing composite literal
    // {{range .}}
      // <li>{{.}}</li>
    // {{end}}
  // range and end to display the data passed
  sages := []string {"Ghandi", "Buddha", "Jesus"}
  err = tpl.ExecuteTemplate(os.Stdout, "test3.gohtml", sages)
  if err != nil {
    log.Fatalln(err)
  }

  example1 := map[string]string{
    "Japan": "Sushi",
    "China": "Noodle",
    "Germany": "Wurst",
    "France": "Pasta",
  }

  err1 := tpl.ExecuteTemplate(os.Stdout, "test3.gohtml", example1)
  if err != nil {
    log.Fatalln(err1)
  }

  germany := property{
    Capital: "Berlin",
    Weather: "Sunny",
  }

  japan := property{
    Capital: "Tokyo",
    Weather: "Rainy",
  }

  france := property{
    Capital: "Paris",
    Weather: "Cloudy",
  }

  country := tpl.ExecuteTemplate(os.Stdout, "test4.gohtml", germany)
  if err != nil {
    log.Fatalln(country)
  }

  countries :=[]property{germany, japan, france}

  example2 := tpl.ExecuteTemplate(os.Stdout, "test5.gohtml", countries)
  if err != nil {
    log.Fatalln(example2)
  }
}

// ExecuteTemplate(writer, file, data to pass)
