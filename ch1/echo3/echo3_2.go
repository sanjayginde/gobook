// Echo3 prints its command-line arguments.
package main

import (
  "fmt"
  "os"
)

func main()  {
  // Exercise 1.2
  for i, arg := range os.Args[1:] {
    fmt.Println(i, arg)
  }
}
