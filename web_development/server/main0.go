package main

import (
  "fmt"
  "io"
  "log"
  "net"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    // Panic is equiv. to Println with call to panic.
    log.Panic(err)
  }
  // defer defers the execution of a function
  // until the surrounding function returns.
  defer li.Close()

  for {
      // Listen => Accept, if successful we get a connection
      conn, err := li.Accept()
      if err != nil {
        // Println formats using the default formats for its operands and writes to standard output.
        // Spaces always added bet. operands and newline is appended.
        log.Println(err)
      }
      // WriteString writes the content of the string s to w, which accepts a slice of bytes.
      io.WriteString(conn, "\nHello from TCP server\n")
      // Fprintln formats using the default formats for its operands and
      // writes to the writer. Spaces are always added between operands and
      // a new line is appended. It returns the number of bytes written and any
      // error encountered.
      fmt.Fprintln(conn, "How is your day?")
      // Fprintf formatis according to a format specifier and writes to
      fmt.Fprintf(conn, "%v", "Well, I hope!")
      // close the connection
      conn.Close()
  }
}
