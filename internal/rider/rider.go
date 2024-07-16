package rider

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/internal/driver"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

var RiderList []*models.Rider
var RiderToDriverMap = make((map[string][]string))

const maxDistance = 5

func AddRider(riderId, x, y string) {
	xCoord, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Invalid x coordinate:", err)
		return
	}
	yCoord, err := strconv.Atoi(y)
	if err != nil {
		fmt.Println("Invalid y coordinate:", err)
		return
	}

	newRider := models.NewRider(riderId, xCoord, yCoord)
	RiderList = append(RiderList, newRider)
}

func MatchRiders(riderId string) {
	var matchedRider *models.Rider
	var matchingDrivers []models.DriverDistanceFromRider
	var matchedDriverIds []string

	for _, rider := range RiderList {
		if rider.RiderId == riderId {
			matchedRider = rider
			break
		}
	}

	for _, driver := range driver.DriverList {
		distance := common.CalculateDistance(driver.DriverXCoord, driver.DriverYCoord, matchedRider.RiderXCoord, matchedRider.RiderYCoord)
		if distance <= maxDistance {
			matchingDrivers = append(matchingDrivers, models.DriverDistanceFromRider{Driver: driver, Distance: distance})
		}
	}

	sort.Slice(matchingDrivers, func(i, j int) bool {
		if matchingDrivers[i].Distance == matchingDrivers[j].Distance {
			return matchingDrivers[i].Driver.DriverId < matchingDrivers[j].Driver.DriverId
		}
		return matchingDrivers[i].Distance < matchingDrivers[j].Distance
	})

	if len(matchingDrivers) == 0 {
		fmt.Println("NO_DRIVERS_AVAILABLE")
	} else {
		fmt.Print("DRIVERS_MATCHED")
		for _, driverDetails := range matchingDrivers {
			driverId := driverDetails.Driver.DriverId
			fmt.Printf(" %s", driverId)
			matchedDriverIds = append(matchedDriverIds, driverId)
		}
		fmt.Println()
		// Store matched driver IDs for the rider
		RiderToDriverMap[riderId] = matchedDriverIds
	}
}
