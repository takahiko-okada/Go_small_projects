# Golang Notes

[Pointers](#pointers)</br>
[Structs](#structs)</br>
&nbsp;&nbsp;&nbsp;[Assigning Struct to a variable](#assigning-struct-to-a-variable)</br>
&nbsp;&nbsp;&nbsp;[Pointer to Struct](#pointer-to-struct)</br>
&nbsp;&nbsp;&nbsp;[Struct Literals](#struct-literals)</br>
[Array](#array)</br>
[Slice](#slice)</br>
[Functions](#functions)</br>
&nbsp;&nbsp;&nbsp;[Arguments and Return Value Type](#arguments-and-return-value-type)</br>
&nbsp;&nbsp;&nbsp;[Named Return Values](#named-return-values)</br>
[Variables](#variables)</br>
&nbsp;&nbsp;&nbsp;[Variable declaration](#variable-declaration)</br>
[Constants](#constants)</br>
[Types](#types)</br>
&nbsp;&nbsp;&nbsp;[Basic types](#basic-types)</br>
&nbsp;&nbsp;&nbsp;[Figuring out variable type](#figuring-out-variable-type)</br>
[Type Conversion](#type-conversion)</br>
&nbsp;&nbsp;&nbsp;[Int to String](#int-to-string)</br>
[Conventions](#conventions)</br>
[Function Naming](#function-naming)</br>
[Loop](#loop)</br>
&nbsp;&nbsp;&nbsp;[For loop using init and post statements](#for-loop-using-init-and-post-statements)</br>
&nbsp;&nbsp;&nbsp;[For loop without init and post statements](#for-loop-without-init-and-post-statements)</br>
&nbsp;&nbsp;&nbsp;[While loop in Go](#while-loop-in-go)</br>
&nbsp;&nbsp;&nbsp;[Infinite loop](#infinite-loop)</br>


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

### Assigning Struct to a variable

```go
package main

import "fmt"

func main() {
  p1 := Person{}                   // Empty struct assign Name: "", Age: 0 by default
  p1.Name = "James Bond"
  p1.Age = 42
  fmt.Println(p1.Name, p1.Age)
}
```

### Pointer to Struct

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

### Struct Literals

```go
package main

import "fmt"

type Location struct {
  lat, lng int
}

var (                     // Defining variables outside of func main
  v1 = Location{50, 100}  // has type Location
  v2 = Location{X: 70}    // Y:0 is implicit
  p  = &Location{20, 90}  // has type *Location
)

func main() {
  fmt.Println(v1, p, v2)
}

```

### Array
#### Array size cannot be changed later, whereas Slice allows this.

```go
package main

import "fmt"

func main() {
  var a [3]string            // initiates an array of 3 strings
  a[0] = "USA"               // push a values
  a[1] = "Japan"
  a[2] = "Germany"

  fmt.Println(a)             // the array itself can be printed
  fmt.Println(a[1])          // prints a specified item
}
```

### Slice
#### Slices works as references to arrays
#### A slice does not store any data, it shows a part of an array
#### Modification to a slice will be reflected to its underlying array

```go
package main

import "fmt"

func main() {
  nations := [3]string{
    "USA",
    "Japan",
    "Germany",                  // Need a comma for the last item
  }                             // Defined an array nations

  s1 := nations[0:1]            // Slices created
  s2 := nations[1:2]

  fmt.Println(s1, s2, nations)
}
```

## Functions
### Arguments and Return Value Type
```go
package main

import "fmt"

func add(x int, y int) int {                     // returns a value, single return value type
  return x + y
}

func plusAndMinus(x int, y int) (int, int) {     // returns two values, two return value type
  return x + y, x - y
}

func main() {
  fmt.Println(add(222, 777))
  fmt.Println(plusAndMinus(222, 777))
}
```

### Named Return Values

```go
package main

import "fmt"

func split(total int) (x, y float64) {
  x = float64(total) * 0.4            // int must be converted before multiplying it by a float
  y = float64(total) * 0.6
  return                              // return without argument returns the Named Return Values x, y
}

func main() {
  fmt.Println(split(100))
}
```

## Variables
### Variable declaration

```go
package main

import "fmt"

var a, b bool
var c, d = 1, "Yes"                  // Declaration With initializer(value), type can be omitted
g := "Fish"                          // Short variable declaration

func main() {
  var e, f bool                      // Variable declaration can be done within functions
  fmt.Println(a, b, c, d, e, f)
}
```

## Constants

```go
package main

import "fmt"

const Pi = 3.14                      // Declared with a capital letter
                                     // No shorthand available

func main() {
  const World = "Die Welt"
  fmt.Println("Guten tag", World)
}
```

## Types
### Basic types

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

### Figuring out variable type
```go
package main

import "fmt"

func main() {
  v := 100
  fmt.Printf("v is of type %T\n", v)           // Printf(format string, a...interface{})(n int, err error)
}
```

### Type Conversion
#### Int to String
```go
package main

import (
  "fmt"
  "strconv"                       // Use "strconv" package
)

func main() {
  i := 5
  s := strconv.Itoa(i)            // Then use "Itoa" function
  fmt.Println(s)
}
```

## Conventions
### Function Naming
```go
// func only available within the package

func internalUseOnly() {            // Lower Camel Case
}

// func to be exported

func ExternalUsePossible() {        // Upper Camel Case
}
```

## Loop
### For loop using init and post statements
```go
package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {         // initial i is 0, i increases each time
    sum += i
  }
  fmt.Println(sum)
}

```

### For loop without init and post statements
```go
package main

import "fmt"

func main() {
  sum := 1
  for ; sum < 10; {
    sum += sum                      // 2 => 4 => 8 => 16, loop ends when sum >= 10
  }
  fmt.Println(sum)
}
```

### While loop (in Go)
```go
func main() {
  sum := 0
  for sum < 10 {                    // +1 each time to sum until 10
    sum += 1
  }
  fmt.Println(sum)
}
```

### Infinite loop

```go
package main

func main() {
  for {
  }
}
```

## Interface

```go
package main

import "fmt"

type German struct{}

func (g *German) Hello() {
  fmt.Println("Guten tag")
}

type English struct{}

func (a *English) Hello() {
  fmt.Println("Hello")
}

func helloInGerman(g *German) {
  g.Hello()
}

func helloInEnglish(a *English) {
  a.Hello()
}

func main() {
  german := new(German)
  english := new(English)
  helloInGerman(german)             // a lot of duplications
  helloInEnglish(english)
}
```

```go
package main

import "fmt"

type Language interface {       // Language Interface
  greeting()
}

type German struct{}

func (g *German) greeting() {
  fmt.Println("Guten tag")
}

type English struct{}

func (c *English) greeting() {
  fmt.Println("Hello")
}

func sayHello(l Language) {    // sayHello takes the interface, solves the duplication
  l.greeting()
}

func main() {
  german := new(German)
  english := new(English)
  sayHello(german)
  sayHello(english)
}
```

## Errors
Go does not have exceptions, errors have to be handled manually
### Check if there's any error while running using console
```shell
errcheck       // returns error messages if there's any
```

## Testing
Create a file for testing in the following manner
```shell
FILENAME_test.go        // insert "_test" after the file name
```
File name with _test will be ignored by
```shell
go get
go install
go build
```
