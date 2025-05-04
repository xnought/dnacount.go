package main

import (
	"fmt"
	"os"
)

func count_frequencies(data []byte) map[byte]uint64 {
	char_count := map[byte]uint64{} // char -> count
	for _, d := range data {
		_, key_exists := char_count[d]
		if key_exists {
			char_count[d]++
		} else {
			char_count[d] = 1
		}
	}
	return char_count
}

func main() {
	data, err := os.ReadFile("main.go")
	if err != nil {
		panic(err)
	}
	out := count_frequencies(data)
	for k, v in range out {
		fmt.Printf("%c=%d\n", k, v);
	}
}
