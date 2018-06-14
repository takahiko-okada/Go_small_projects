# Golang Notes

## Pointers
### Pointers holds the memory address of the value

```go
package main

import "fmt"

func main() {
  i := 42          // assign values to the variables
  p := &i          // p is the memory address of i
  fmt.Println(p)   // returns the address
  fmt.Println(*p)  // returns the value of i, * operator shows the value the
                   // address holds

  *p = 21          // re-assigns the value to i, through * operator
  fmt.Println(p)   // the memory address is the same as before
  fmt.Println(*p)  // shows the re-assigned value
}

```
## Structs
### Struct is a collection of fields.

```go
package main

import "fmt"

type Person struct {
  Name string      // field name, with Capital letter and value type
  Age int
}

func main() {
  fmt.Println(Person{"James Bond", 42}  // Prints a struct {"James Bond" 42}
}

```


