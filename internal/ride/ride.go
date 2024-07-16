package ride

import (
	"fmt"
	"strconv"

	"github.com/PAlagusurya/geektrust/internal/driver"
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
	var matchedDrivers []string
	nthDriver, err := strconv.Atoi(nthDriverIndex)

	if err != nil {
		fmt.Printf("Error converting '%s' to integer: %v\n", nthDriverIndex, err)
	}

	//You cannot start the ride which is already started
	for _, ride := range RideList {
		if ride.RideId == rideId {
			fmt.Println("INVALID_RIDE")
			return
		}
	}

	if matchedDriverIds, exists := rider.RiderToDriverMap[riderId]; exists {
		matchedDrivers = matchedDriverIds
		if nthDriver > len(matchedDriverIds) {
			fmt.Println("DRIVER NOT AVAILABLE")
			return
		}
	}

	requestedDriver := matchedDrivers[nthDriver-1]

	for _, driver := range driver.DriverList {
		if driver.DriverId == requestedDriver {
			if driver.IsAvailable {
				ride := models.RideDetails{
					RideId:   rideId,
					RiderId:  riderId,
					DriverId: driver.DriverId,
				}
				RideList = append(RideList, &ride)
				driver.IsAvailable = false
				fmt.Println("RIDE_STARTED" + " " + rideId)
			} else {
				fmt.Println("DRIVER NOT AVAILABLE")
				return
			}
		}
	}
}
