package ride

import (
	"fmt"
	"math"

	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/internal/driver"
	"github.com/PAlagusurya/geektrust/internal/rider"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

func isRideAlreadyStarted(rideId string) bool {
	for _, ride := range RideList {
		if ride.RideId == rideId {
			return true
		}
	}
	return false
}

func assignDriverToRide(driverId, rideId, riderId string) bool {
	for _, driver := range driver.DriverList {
		if driver.DriverId == driverId && driver.IsAvailable {
			ride := models.RideDetails{
				RideId:          rideId,
				RiderId:         riderId,
				DriverId:        driver.DriverId,
				IsRideCompleted: false,
			}
			RideList = append(RideList, &ride)
			driver.IsAvailable = false
			fmt.Println(common.MessageRideStarted + " " + rideId)
			return true
		}
	}
	return false
}

func CompleteRide(ride *models.RideDetails, xCoord, yCoord, totalTime int) {
	for _, rider := range rider.RiderList {
		if rider.RiderId == ride.RiderId {
			totalDistanceTravelled := common.CalculateDistance(rider.RiderXCoord, rider.RiderYCoord, xCoord, yCoord)
			ride.Distance = totalDistanceTravelled
			ride.TotalTimeTaken = totalTime
			ride.IsRideCompleted = true
			return
		}
	}
}

func CalculateBillAmount(distance float64, totalTime int) float64 {
	preTaxAmount := distance*common.AdditionalFarePerKm + float64(totalTime)*common.AdditionalFarePerMin + common.BaseFare
	billAmount := preTaxAmount + (preTaxAmount * common.ServiceTaxRate)
	return math.Round(billAmount*common.DistancePrecision) / common.DistancePrecision
}
