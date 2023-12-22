package services

import (
	"travel_booking/databases"
	"travel_booking/typ"
)

func CreateFlight(flight *typ.CreateFlightReq) error {
	return databases.CreateFlight(flight)
}

func DeleteFlight(flight *typ.DeleteFlightReq) error {
	return databases.DeleteFlight(flight)
}

func UpdateFlight(flight *typ.UpdateFlightReq) error {
	return databases.UpdateFlight(flight)
}

func GetFlight() (typ.GetFlightRes, error) {
	return databases.GetFlight()
}
