package main

import (
	"fmt"
	"os"
)

type Option string

const (
	Create Option = "create"
	Delete Option = "delete"
	Rename Option = "rename"
)

func main() {
	helpMessage := "Usage: command <option> <arguments>"

	if len(os.Args) < 1 {
		fmt.Println(helpMessage)
		return
	}

	option := Option(os.Args[1])

	switch option {
		case Create:
			path := os.Args[2]
			_, err := os.Create(path)
			if err != nil {
				printError(err.Error())
				return
			}
		case Delete:
			path := os.Args[2]
			err := os.Remove(path)
			if err != nil {
				printError(err.Error())
				return
			}
		case Rename:
			oldPath := os.Args[2]
			newPath := os.Args[3]
			err := os.Rename(oldPath, newPath)
			if err != nil {
				printError(err.Error())
				return
			}
		default:
			fmt.Println(helpMessage)
			return
	}
}

func printError(error string) {
	fmt.Println("An error has occurred:", error)
}
