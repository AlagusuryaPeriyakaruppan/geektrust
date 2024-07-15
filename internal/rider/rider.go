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
var MatchesMap = make(map[string][]string)

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

	rider := models.NewRider(riderId, xCoord, yCoord)
	RiderList = append(RiderList, rider)
}

func MatchRiders(riderId string) {
	var matchedRider *models.Rider
	var driverDistance []models.DriverDistanceFromRider
	var matchedDriverIds []string

	for _, rider := range RiderList {
		if rider.RiderId == riderId {
			matchedRider = rider
			break
		}
	}

	if matchedRider == nil {
		fmt.Println("Rider not found")
		return
	}

	for _, driver := range driver.DriverList {
		distance := common.CalculateDistance(driver.DriverXCoord, driver.DriverYCoord, matchedRider.RiderXCoord, matchedRider.RiderYCoord)
		if distance <= 5 {
			driverDistance = append(driverDistance, models.DriverDistanceFromRider{Driver: driver, Distance: distance})
		}
	}

	sort.Slice(driverDistance, func(i, j int) bool {
		if driverDistance[i].Distance == driverDistance[j].Distance {
			return driverDistance[i].Driver.DriverId < driverDistance[j].Driver.DriverId
		}
		return driverDistance[i].Distance < driverDistance[j].Distance
	})

	if len(driverDistance) == 0 {
		fmt.Println("NO_DRIVERS_AVAILABLE")
	} else {
		fmt.Print("DRIVERS_MATCHED")
		for _, driverDetails := range driverDistance {
			driverId := driverDetails.Driver.DriverId
			fmt.Printf(" %s", driverId)
			matchedDriverIds = append(matchedDriverIds, driverId)
		}
		fmt.Println()
		MatchesMap[riderId] = matchedDriverIds
	}
}
