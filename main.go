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
		fmt.Println(argList)

		if len(argList) == 0 {
			fmt.Println("Empty command")
			return
		}

		action := argList[0]
		switch action {
		case "ADD_DRIVER":
			if len(args) != 4 {
				fmt.Println("Invalid number of arguments for ADD_DRIVER")
				return
			}
			driver.AddDriver(argList[1], argList[2], argList[3])
		case "ADD_RIDER":
			if len(args) != 4 {
				fmt.Println("Invalid number of arguments for ADD_RIDER")
				return
			}
			rider.AddRider(argList[1], argList[2], argList[3])
		case "MATCH":
			if len(args) != 2 {
				fmt.Println("Invalid number of arguments for MATCH")
				return
			}
			rider.MatchRiders(argList[1])
		case "START_RIDE":
			if len(args) != 4 {
				fmt.Println("Invalid number of arguments for START_RIDE")
				return
			}
			ride.StartRide(argList[1], argList[2], argList[3])
		default:
			fmt.Println("ACTION NOT FOUND")
		}

	}
}
