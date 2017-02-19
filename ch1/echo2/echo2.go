// Echo2 prints its command-line arguments.
package main

import (
  "fmt"
  "os"
)

func main()  {
  result, sep := "", ""
  for _, arg := range os.Args[1:] {
    result += sep + arg
    sep = " "
  }
  fmt.Println(result)
}
