package driver

import (
	"fmt"
	"strconv"

	"github.com/PAlagusurya/geektrust/pkg/models"
)

var DriverList []*models.Driver

func AddDriver(driverId, x, y string) {
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

	driver := models.NewDriver(driverId, xCoord, yCoord)
	DriverList = append(DriverList, driver)
}
