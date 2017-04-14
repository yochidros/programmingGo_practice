package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	number := os.Args[1]
	fmt.Println(comma(number))
	extensionComma(number)
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func extensionComma(s string) string {
	var buf bytes.Buffer
	separateCount := len(s) / 3
	buf.WriteString("fjdk")
	for i, v := range s {
		fmt.Println(i, string(v))
	}
	fmt.Println(separateCount)
	return s
}
