package rider

import (
	"fmt"
	"sort"

	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/internal/driver"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

func FindRiderById(riderId string) *models.Rider {
	for _, rider := range RiderList {
		if rider.RiderId == riderId {
			return rider
		}
	}
	return nil
}

func FindMatchingDrivers(matchedRider *models.Rider) []models.DriverDistanceFromRider {
	var matchingDrivers []models.DriverDistanceFromRider

	for _, driver := range driver.DriverList {
		distance := common.CalculateDistance(driver.DriverXCoord, driver.DriverYCoord, matchedRider.RiderXCoord, matchedRider.RiderYCoord)
		if distance <= maxDistance && driver.IsAvailable {
			matchingDrivers = append(matchingDrivers, models.DriverDistanceFromRider{Driver: driver, Distance: distance})
		}
	}
	return matchingDrivers
}

func SortAndPrintMatchingDrivers(matchingDrivers []models.DriverDistanceFromRider) {
	sort.Slice(matchingDrivers, func(i, j int) bool {
		if matchingDrivers[i].Distance == matchingDrivers[j].Distance {
			return matchingDrivers[i].Driver.DriverId < matchingDrivers[j].Driver.DriverId
		}
		return matchingDrivers[i].Distance < matchingDrivers[j].Distance
	})

	fmt.Print(common.MessageDriversMatched)
	for _, driverDetails := range matchingDrivers {
		fmt.Printf(" %s", driverDetails.Driver.DriverId)
	}
	fmt.Println()
}

func StoreMatchedDrivers(riderId string, matchingDrivers []models.DriverDistanceFromRider) {
	var matchedDriverIds []string
	for _, driverDetails := range matchingDrivers {
		matchedDriverIds = append(matchedDriverIds, driverDetails.Driver.DriverId)
	}
	RiderToDriverMap[riderId] = matchedDriverIds
}
