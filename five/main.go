package main
 
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
 
func palindrome(input string, n int) bool {
    if n == 0 || n == 1 {
      return true
    } else if input[0] == input[n-1] {
      return palindrome(input[1:n-1], n-2)
    } else {
      return false
    }
}

func main() {
    consoleReader := bufio.NewReader(os.Stdin)
    fmt.Println("Input some string :  ")
    input, _ := consoleReader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    var n int = len(input)
    fmt.Println(palindrome(input, n))
}
