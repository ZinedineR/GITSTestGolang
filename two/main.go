package main

import (
    "fmt"
    "strings"
    "regexp"
)

func main() {
    var email string
    pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    fmt.Println("Input your email : ")
    fmt.Scan(&email)
    matches := pattern.MatchString(email)
    at := strings.LastIndex(email, "@")
    period := strings.Index(email, ".")
    if at >= 1 && period >=1 && matches == true{
        username, domain := email[:at], email[period:]
        var length = len([]rune(username))
        if length <20 {
          if domain == ".id" || domain == ".co.id"{    
            fmt.Printf("Username: %s, Domain: %s\nThis email is valid", username, domain)
          } else {
          fmt.Printf("Error: %s is an invalid email address, only .co.id and .id domains allowed\n", email)
          }
        } else {
        fmt.Printf("Error: %s is an invalid email address only 20 characters max allowed before @\n", email)
        }
    } else {
        fmt.Printf("Error: %s is an invalid email address, please make sure there's @ and using co.id/.id domain\n", email)
    }
}