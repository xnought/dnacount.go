package main

import (
	"fmt"
	"os"
)

func main()  {
	data, err := os.ReadFile("main.go")
	if err != nil {
		panic(err)
	}
	for i, d := range data {
		if d == '\n' {
			fmt.Printf("%d: '\\n'\n", i)
		} else if d == '\t' {
			fmt.Printf("%d: '\\t'\n", i)
		} else {
			fmt.Printf("%d: '%c'\n", i, d)
		}
	}
}