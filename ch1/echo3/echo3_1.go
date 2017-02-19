// Echo3 prints its command-line arguments, incl. the command
package main

import (
  "fmt"
  "os"
  "strings"
)

func main()  {
  // Exercise 1.1
  fmt.Println(strings.Join(os.Args[0:], " "))
}
