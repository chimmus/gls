package main

import ( 
	"fmt"
	"os"
	"strings"
)

// 15:11

func main() {
	path := "."
	
	show_all := false

	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "-") {
			path = arg
			continue
		}
		for _, char := range arg[1:] {
			switch char {
			case 'a':
				show_all = true
			default:
				fmt.Fprintf(os.Stderr, "usage: %s [-a] [path]\n", os.Args[0])
				os.Exit(1)
			}
		}
	}
	entries, error := os.ReadDir(path)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		os.Exit(1)
	}
	for _, entries := range entries {
		if !show_all && strings.HasPrefix(entries.Name(), ".") {
			continue
		}
		fmt.Println(entries.Name())
	}
}
