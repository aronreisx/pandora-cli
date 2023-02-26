package app

import (
	"fmt"
	"os"
)

var HelpMessage = "Usage: senses <option> <arguments>"

func Error(error string) {
	fmt.Println(error)
	os.Exit(1)
}

func Success() {
	os.Exit(0)
}

func MinimumArgsAmountValidator(minimumArgsNumber int) {
	if len(os.Args) < minimumArgsNumber {
		Error(HelpMessage)
	}
}