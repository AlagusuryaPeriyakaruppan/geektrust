package common

import "errors"

// Error constants
var (
	ErrorInvalidArguments = errors.New("Invalid arguments")
	ErrorInvalidAction    = errors.New("Invalid action")
)

// Command constants
const (
	CommandAddDriver = "ADD_DRIVER"
	CommandAddRider  = "ADD_RIDER"
	CommandMatch     = "MATCH"
	CommandStartRide = "START_RIDE"
	CommandStopRide  = "STOP_RIDE"
	CommandBill      = "BILL"
)

// Error and message constants
const (
	ErrorInvalidXCoord        = "Invalid x coordinate:"
	ErrorInvalidYCoord        = "Invalid y coordinate:"
	ErrorInvalidTime          = "Invalid time format"
	ErrorInvalidRide          = "INVALID_RIDE"
	ErrorTimeTaken            = "INVALID_TIME_TAKEN"
	MessageRideStopped        = "RIDE_STOPPED"
	MessageRideStarted        = "RIDE_STARTED"
	ErrorDriverNotAvailable   = "DRIVER NOT AVAILABLE"
	MessageNoDriversAvailable = "NO_DRIVERS_AVAILABLE"
	MessageDriversMatched     = "DRIVERS_MATCHED"
	MessageRideNotCompleted   = "RIDE_NOT_COMPLETED"
)

// Fare-related constants
const (
	BaseFare             = 50.0
	AdditionalFarePerKm  = 6.5
	AdditionalFarePerMin = 2.0
	ServiceTaxRate       = 0.20
	DistancePrecision    = 100
)
