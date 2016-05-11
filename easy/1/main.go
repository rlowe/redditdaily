package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  var username string
  var age int

  fmt.Print("Enter your name: ")
  bio := bufio.NewReader(os.Stdin)
  name, _, _ := bio.ReadLine() // We will assume that it isn't _that_ long
  fmt.Print("Enter your age (in years): ")
  fmt.Scanln(&age)
  fmt.Print("Enter your reddit username: ")
  fmt.Scanln(&username) // Scanln because Reddit usernames cannot have spaces :-D

  fmt.Printf("Your name is %s, you are %d years old, and your username is %s\n", string(name), age, username)
}
