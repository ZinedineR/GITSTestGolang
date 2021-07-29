package main

import (
  "fmt"
)

func main() {
  var input int
  var save int
  var save1 int
  fmt.Println("Input your number: ")
  fmt.Scanln(&input)
  fmt.Printf("input => %d \n", input)
  save = input % 3
  save1 = input % 5
  if save1 == 0 && save == 0 {
	fmt.Println("Hello World")
	} else if save1 == 0 {
	fmt.Println("World")
	} else if save == 0 {
	fmt.Println("Hello")
	} else {
  fmt.Println("Sorry, please input another number")
  }
}