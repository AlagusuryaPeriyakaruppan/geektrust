package driver

import (
	"github.com/PAlagusurya/geektrust/internal/common"
	"github.com/PAlagusurya/geektrust/pkg/models"
)

var DriverList []*models.Driver

func AddDriver(driverId, x, y string) {
	xCoord, err := common.ConvertStringToInt(x, common.ErrorInvalidXCoord)
	if err != nil {
		return
	}

	yCoord, err := common.ConvertStringToInt(y, common.ErrorInvalidYCoord)
	if err != nil {
		return
	}
	newDriver := models.NewDriver(driverId, xCoord, yCoord)
	DriverList = append(DriverList, newDriver)
}
