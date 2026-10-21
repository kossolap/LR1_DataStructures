package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "file.data"
	query := ""

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--file":
			if i+1 < len(os.Args) {
				i++
				filename = os.Args[i]
			}
		case "--query":
			if i+1 < len(os.Args) {
				i++
				query = os.Args[i]
			}
		default:
			fmt.Fprintf(os.Stderr, "error: unknown argument %s\n", os.Args[i])
			os.Exit(1)
		}
	}

	db := NewDatabase()
	db.LoadFromFile(filename)

	if query != "" {
		ok, result := db.ExecuteCmd(query)
		if result != "" {
			fmt.Println(result)
		}
		if ok && isModified(query) {
			db.SaveToFile(filename)
		}
	} else {
		fmt.Println("no query provided. use --query \"COMMAND\".")
	}
}

func isModified(query string) bool {
	for _, sub := range []string{"PUSH", "DEL", "POP", "INSERT", "SET"} {
		if contains(query, sub) {
			return true
		}
	}
	return false
}

func contains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
