package models

type Driver struct {
	DriverId     string
	DriverXCoord int
	DriverYCoord int
	IsAvailable  bool
}

func NewDriver(driverId string, x, y int) *Driver {
	return &Driver{
		DriverId:     driverId,
		DriverXCoord: x,
		DriverYCoord: y,
		IsAvailable:  true,
	}
}
