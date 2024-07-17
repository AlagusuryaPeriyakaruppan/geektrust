package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PAlagusurya/geektrust/internal/inputhandler"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")
		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputhandler.ProcessInput(scanner.Text())
	}

}
