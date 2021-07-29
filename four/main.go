package main
 
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
 
func reverse(str string) (result string) {
    for _, v := range str {
        result = string(v) + result
    }
    return
}

func main() {
    consoleReader := bufio.NewReader(os.Stdin)
    fmt.Println("Input some string :  ")
    input, _ := consoleReader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")
    parsed := reverse(input)
    fmt.Println(input, "=>", parsed)
}
