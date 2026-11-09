package main

import ( 
	"fmt"
	"os"
)

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	entries, _ := os.ReadDir(path)
	for _, entries := range entries {
		fmt.Println(entries.Name())
	}
}
