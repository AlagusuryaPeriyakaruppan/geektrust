package models

type Rider struct {
	RiderId     string
	RiderXCoord int
	RiderYCoord int
}

func NewRider(riderId string, x, y int) *Rider {
	return &Rider{
		RiderId:     riderId,
		RiderXCoord: x,
		RiderYCoord: y,
	}
}
