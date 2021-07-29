package main

import (
  "fmt"
  "os"
)

func main() {
  var hour, minute, second int
  var pm string
  fmt.Println("Input hour (12 hours format:")
  fmt.Scanln(&hour)
  if hour > 12 {
    fmt.Println("Error, please input again")
    os.Exit(1)
  }
  fmt.Println("Input minute (0-59):")
  fmt.Scanln(&minute)
  if minute > 59 {
    fmt.Println("Error, please input again")
    os.Exit(1)
  }
  fmt.Println("Input second (0-59):")
  fmt.Scanln(&second)
  if second > 59 {
    fmt.Println("Error, please input again")
    os.Exit(1)
  }
  fmt.Println("please type PM or AM ? :")
  fmt.Scanf("%s", &pm)
  if pm == "PM" || pm == "pm" {
    if hour != 12 {
      hour += 12
      fmt.Println(hour,":",minute,":",second)
    } else {
      fmt.Println(hour,":",minute,":",second)
    }
  } else if pm == "AM" || pm == "am" {
     if hour == 12 {
      hour = 00
      fmt.Println(hour,":",minute,":",second)
    } else {
      fmt.Println(hour,":",minute,":",second)
    }
  }
}