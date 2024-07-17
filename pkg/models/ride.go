package models

type RideDetails struct {
	RideId          string
	RiderId         string
	DriverId        string
	IsRideCompleted bool
	Distance        float64
	Bill            float64
	TotalTimeTaken  int
}

type DriverDistanceFromRider struct {
	Driver   *Driver
	Distance float64
}
