package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
)

func main() {
  xys, err := readData("DataScience/data.txt")
  if err != nil {
    log.Fatalf("could not read data.txt: %v", err)
  }
  for _, xy := range xys {
    fmt.Println(xy.x, xy.y)
  }
}

type xy struct { x, y float64 }

func readData(path string) ([]xy, error){
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  var xys []xy
  s := bufio.NewScanner(f)
  for s.Scan() {
      var x, y float64
      _, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y) // only err is needed to be assigned
      if err != nil {
        log.Printf("Discarding bad data %q: %v", s.Text(), err)
      }
      xys = append(xys, xy{x, y})
  }
  if err := s.Err(); err != nil {
    return nil, fmt.Errorf("could not scan: %v")
  }
  return xys, nil
}
