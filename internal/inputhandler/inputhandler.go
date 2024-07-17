package inputhandler

import (
	"strings"

	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/internal/driver"
	"github.com/PAlagusurya/geektrust/internal/ride"
	"github.com/PAlagusurya/geektrust/internal/rider"
)

func ProcessInput(args string) {
	argList := strings.Fields(args)
	if len(argList) == 0 {
		return
	}

	action := argList[0]

	if !common.ValidateArguments(action, argList) {
		common.HandleError(common.ErrorInvalidArguments)
		return
	}

	switch action {
	case common.CommandAddDriver:
		driver.AddDriver(argList[1], argList[2], argList[3])
	case common.CommandAddRider:
		rider.AddRider(argList[1], argList[2], argList[3])
	case common.CommandMatch:
		rider.MatchRiders(argList[1])
	case common.CommandStartRide:
		ride.StartRide(argList[1], argList[2], argList[3])
	case common.CommandStopRide:
		ride.StopRide(argList[1], argList[2], argList[3], argList[4])
	case common.CommandBill:
		ride.CalculateBill(argList[1])
	default:
		common.HandleError(common.ErrorInvalidAction)
	}
}
