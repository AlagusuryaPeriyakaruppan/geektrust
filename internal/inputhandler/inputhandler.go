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

	switch action {
	case common.CommandAddDriver:
		if len(argList) < 4 {
			common.HandleError(common.ErrorInvalidArguments)
			return
		}
		driver.AddDriver(argList[1], argList[2], argList[3])
	case common.CommandAddRider:
		if len(argList) < 4 {
			common.HandleError(common.ErrorInvalidArguments)
			return
		}
		rider.AddRider(argList[1], argList[2], argList[3])
	case common.CommandMatch:
		if len(argList) < 2 {
			common.HandleError(common.ErrorInvalidArguments)
			return
		}
		rider.MatchRiders(argList[1])
	case common.CommandStartRide:
		if len(argList) < 4 {
			common.HandleError(common.ErrorInvalidArguments)
			return
		}
		ride.StartRide(argList[1], argList[2], argList[3])
	case common.CommandStopRide:
		if len(argList) < 5 {
			common.HandleError(common.ErrorInvalidArguments)
			return
		}
		ride.StopRide(argList[1], argList[2], argList[3], argList[4])
	case common.CommandBill:
		if len(argList) < 2 {
			common.HandleError(common.ErrorInvalidArguments)
			return
		}
		ride.CalculateBill(argList[1])
	default:
		common.HandleError(common.ErrorInvalidAction)
	}
}
