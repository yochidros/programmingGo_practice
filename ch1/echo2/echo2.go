package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	// printCommandLine()
	simplePrintCommandLine()
	end := time.Since(start).Seconds()
	fmt.Println(end)
}
func simplePrintCommandLine() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
func printCommandLine() {
	var sep string
	for i, arg := range os.Args[0:] {
		sep = " "
		fmt.Println(i, sep, arg)
	}

}
