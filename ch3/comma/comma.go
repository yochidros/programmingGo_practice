package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	number := os.Args[1]
	fmt.Printf("recursive comma: %s\n", comma(number))
	fmt.Printf("non recursive comma: %s\n", extensionComma(number))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

/* practice 3.5_10 */
func extensionComma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	start := n % 3
	if start == 0 {
		start = 3
	}
	for i, v := range s {
		fmt.Fprintf(&buf, "%c", v)
		start--
		if start == 0 && i < n-1 {
			buf.WriteByte(',')
			start = 3
		}
	}
	return buf.String()
}
