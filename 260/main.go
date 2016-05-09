package main

import (
  "fmt"
  "garage"
)

func main() {
  d := garage.Door{}
  err := d.Click("button_clicked")
  if err != nil {
    fmt.Println("ERR", err)
  }
}

