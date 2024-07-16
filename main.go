package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PAlagusurya/geektrust/internal/driver"
	"github.com/PAlagusurya/geektrust/internal/ride"
	"github.com/PAlagusurya/geektrust/internal/rider"
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
		fmt.Println("Error opening the input file")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)

		action := argList[0]

		switch action {
		case "ADD_DRIVER":
			driver.AddDriver(argList[1], argList[2], argList[3])
		case "ADD_RIDER":
			rider.AddRider(argList[1], argList[2], argList[3])
		case "MATCH":
			rider.MatchRiders(argList[1])
		case "START_RIDE":
			ride.StartRide(argList[1], argList[2], argList[3])
		default:
			fmt.Println("ACTION NOT FOUND")
		}

	}
}
