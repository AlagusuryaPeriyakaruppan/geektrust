package ride

import (
	"fmt"
	"strconv"

	"github.com/PAlagusurya/geektrust/internal/driver"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

var Rides []*models.RideDetails

var MatchesMap = make(map[string][]string)

func StartRide(rideId, n, riderId string) {
	nthValue, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("Invalid value for n:", err)
		return
	}

	if matchedDriverIds, exists := MatchesMap[riderId]; exists {
		if nthValue >= len(matchedDriverIds) {
			fmt.Println("Invalid nth value")
			return
		}
		matchedDriverId := matchedDriverIds[nthValue]
		for _, driver := range driver.DriverList {
			if driver.DriverId == matchedDriverId {
				if !driver.IsAvailable {
					fmt.Println("DRIVER NOT AVAILABLE")
					return
				}
				ride := &models.RideDetails{
					RideId:   rideId,
					RiderId:  riderId,
					DriverId: matchedDriverId,
				}
				Rides = append(Rides, ride)
				driver.IsAvailable = false
				fmt.Println("RIDE_STARTED", rideId)
				return
			}
		}
	} else {
		fmt.Println("No matches found for rider")
	}
}
