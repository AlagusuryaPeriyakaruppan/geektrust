package rider

import (
	"fmt"

	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

var RiderList []*models.Rider
var RiderToDriverMap = make((map[string][]string))

const maxDistance = 5

func AddRider(riderId, x, y string) {

	xCoord, yCoord, err := common.ConvertCoordinates(x, y, common.ErrorInvalidXCoord, common.ErrorInvalidYCoord)
	if err != nil {
		fmt.Println(err)
		return
	}

	newRider := models.NewRider(riderId, xCoord, yCoord)
	RiderList = append(RiderList, newRider)
}

func MatchRiders(riderId string) {
	matchedRider := FindRiderById(riderId)
	if matchedRider == nil {
		fmt.Println(common.ErrorInvalidRide)
		return
	}

	matchingDrivers := FindMatchingDrivers(matchedRider)
	if len(matchingDrivers) == 0 {
		fmt.Println(common.MessageNoDriversAvailable)
		return
	}

	SortAndPrintMatchingDrivers(matchingDrivers)
	StoreMatchedDrivers(riderId, matchingDrivers)

}
