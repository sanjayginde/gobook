// Echo1 prints its command-line arguments.
package main

import (
  "fmt"
  "os"
)

func main()  {
  var result, sep string
  for i :=1; i < len (os.Args); i++ {
    result += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(result)
}
