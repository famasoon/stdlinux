package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("argc =", len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("argv[%d] = %s\n", i, os.Args[i])
	}
	return
}
