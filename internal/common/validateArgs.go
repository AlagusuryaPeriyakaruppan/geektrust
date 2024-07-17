package common

func ValidateArguments(action string, argList []string) bool {
	switch action {
	case CommandAddDriver:
		return len(argList) >= MinArgsAddDriver
	case CommandAddRider:
		return len(argList) >= MinArgsAddRider
	case CommandMatch:
		return len(argList) >= MinArgsMatch
	case CommandStartRide:
		return len(argList) >= MinArgsStartRide
	case CommandStopRide:
		return len(argList) >= MinArgsStopRide
	case CommandBill:
		return len(argList) >= MinArgsBill
	default:
		return false
	}
}
