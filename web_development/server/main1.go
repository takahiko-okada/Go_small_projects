package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln(err)
  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Println(err)
      // Continue statement causes the loop to skip the remainder of its body
      // and immediately retest its condition prior to reiterating.
      // In short, it brings you back to the for loop if the statement is true.
      continue
    }
    go handle(conn)
  }
}

func handle(conn net.Conn) {
  // bufio library's  function.
  // bufio.NewScanner() is calling NewScanner() from bufio package
  // returns a new Scanner to read.
  scanner := bufio.NewScanner(conn)
  // func Scan() works on the scanner (type Scanner) struct, which is scanner in this case.
  // Scan advances the Scanner to the nex token. Which will then be available through
  // the Bytes or Text method.
  for scanner.Scan() {
    // ln is a text line, got from scanner variable.
    ln := scanner.Text()
    // print it
    fmt.Println(ln)
  }
  defer conn.Close()

  fmt.Println("Code got here.")
}
