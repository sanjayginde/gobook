// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files.
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filenames := os.Args[1:]

	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

