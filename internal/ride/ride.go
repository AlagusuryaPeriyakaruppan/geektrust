package ride

import (
	"fmt"
	"strconv"

	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/internal/rider"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

/*
1. Check if rideId is already present in RideList. If it is, print "INVALID_RIDE".
2. Ensure that the value of n is less than or equal to the length of the matched drivers for the corresponding rider.
3. Check if the matched driver is available. If they are, share the rideId otherwise, print "DRIVER NOT AVAILABLE".
*/

var RideList []*models.RideDetails

func StartRide(rideId, nthDriverIndex, riderId string) {
	nthDriver, err := strconv.Atoi(nthDriverIndex)

	if err != nil {
		fmt.Printf("Error converting '%s' to integer: %v\n", nthDriverIndex, err)
	}

	//You cannot start the ride which is already started
	if isRideAlreadyStarted(rideId) {
		fmt.Println(common.ErrorInvalidRide)
		return
	}

	matchedDrivers, exists := rider.RiderToDriverMap[riderId]

	if len(matchedDrivers) == 0 {
		fmt.Println(common.ErrorInvalidRide)
		return
	}

	if !exists || nthDriver > len(matchedDrivers) {
		fmt.Println(common.ErrorDriverNotAvailable)
		return
	}

	requestedDriver := matchedDrivers[nthDriver-1]

	if !assignDriverToRide(requestedDriver, rideId, riderId) {
		fmt.Println(common.ErrorDriverNotAvailable)
	}

}

// StopRide stops a ride and calculates the distance traveled
func StopRide(rideId, x, y, time string) {
	xCoord, yCoord, err := common.ConvertCoordinates(x, y, common.ErrorInvalidXCoord, common.ErrorInvalidYCoord)
	if err != nil {
		common.HandleError((err))
		return
	}

	totalTime, err := common.ConvertStringToInt(time, common.ErrorInvalidTime)
	if err != nil {
		common.HandleError(err)
		return
	}

	for _, ride := range RideList {
		if ride.RideId == rideId && !ride.IsRideCompleted {
			CompleteRide(ride, xCoord, yCoord, totalTime)
			fmt.Printf("%s %s\n", common.MessageRideStopped, rideId)
			return
		}
	}

	fmt.Println(common.ErrorInvalidRide)
}

func CalculateBill(rideId string) {
	ride := FindRideByID(rideId)
	if ride == nil {
		fmt.Println(common.ErrorInvalidRide)
		return
	}

	if !ride.IsRideCompleted {
		fmt.Println(common.MessageRideNotCompleted)
		return
	}

	ride.Bill = CalculateBillAmount(ride.Distance, ride.TotalTimeTaken)
	fmt.Printf("BILL %s %s %0.2f\n", ride.RideId, ride.DriverId, ride.Bill)
}
