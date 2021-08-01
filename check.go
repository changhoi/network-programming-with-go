package main

import (
	"fmt"
	"os"
)

func ArgsCheck(length int) {
	if len(os.Args) != length {
		fmt.Fprintf(os.Stderr, "Usage is wrong")
		os.Exit(1)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
