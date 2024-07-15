package models

type RideDetails struct {
	RideId   string
	DriverId string
	RiderId  string
}

type DriverDistanceFromRider struct {
	Driver   *Driver
	Distance float64
}
