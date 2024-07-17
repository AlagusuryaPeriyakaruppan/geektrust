package common

import "errors"

// Error constants
var (
	ErrorInvalidArguments = errors.New("invalid arguments")
	ErrorInvalidAction    = errors.New("invalid action")
)

// Arguments constants
const (
	MinArgsAddDriver = 4
	MinArgsAddRider  = 4
	MinArgsMatch     = 2
	MinArgsStartRide = 4
	MinArgsStopRide  = 5
	MinArgsBill      = 2
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
