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

#### Assigning Struct to a variable

```go
package main

import "fmt"

func main() {
  p1 := Person{}
  p1.Name = "James Bond"
  p1.Age = 42
  fmt.Println(p1.Name, p1.Age)
}
```

#### Pointer to Struct

```go
package main

import "fmt"

func main() {
  p1 := Person{"James Bond", 42}
  v := &p1
  v.Name = "007"                  // Go allows changing the value of p1
                                  // without doing *v.Name
  fmt.Println(p1)
}
```






