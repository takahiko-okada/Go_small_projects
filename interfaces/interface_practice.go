package main

import (
  "fmt"
  "math"
)

type geometry interface {
  area() float64
  perim() float64
  // perimeter => the continuous line forming the bounary of a closed geometrical figure.
}

type rectangle struct {
  width, height float64
}

type circle struct {
  radius float64
}

func (r rectangle) area() float64 {
  return r.width * r.height
}
