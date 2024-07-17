package driver

import (
	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

var DriverList []*models.Driver

func AddDriver(driverId, x, y string) {
	xCoord, yCoord, err := common.ConvertCoordinates(x, y, common.ErrorInvalidXCoord, common.ErrorInvalidYCoord)
	if err != nil {
		common.HandleError(err)
		return
	}
	newDriver := models.NewDriver(driverId, xCoord, yCoord)
	DriverList = append(DriverList, newDriver)
}
