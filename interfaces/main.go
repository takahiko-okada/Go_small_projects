package main

import "fmt"

// new custom type called bot,
// all the types that has
//  getGreeting() string
// is a member of this bot type!
type bot interface {
  // define customized rules inside interface type
  // getGreeting(int) string
  // things that
  getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
  eb := englishBot{}
  sb := spanishBot{}

  printGreeting(eb)
  printGreeting(sb)
}

func printGreeting(b bot) {
  fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
//   fmt.Println(eb.getGreeting())
// }

// func pringGreeting(sb spanishBot) {
//   fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
  // Custom logic for generating english greeting
  return "Hi there!"
}


func (sb spanishBot) getGreeting() string {
  return "Hola"
}
