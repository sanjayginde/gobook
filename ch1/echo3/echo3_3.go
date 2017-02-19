// Echo3_3 prints benchmarks string concat with command line arguments
package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func main()  {
  // Exercise 1.3
  iterations := 10000
  start := time.Now()
  for i := 0; i < iterations; i++ {
    var result, sep string
    for i :=1; i < len (os.Args); i++ {
      result += sep + os.Args[i]
      sep = " "
    }
    fmt.Println(result)
  }
  string_concat_elapsed := time.Since(start).Nanoseconds()


  start = time.Now()
  for i := 0; i < iterations; i++ {
    fmt.Println(strings.Join(os.Args[1:], " "))
  }
  strings_join_elapsed := time.Since(start).Nanoseconds()

  fmt.Println("string concat: ", string_concat_elapsed)
  fmt.Println("strings.Join:  ", strings_join_elapsed)
  fmt.Println("diff:          ", string_concat_elapsed - strings_join_elapsed)
}
