package services

import (
	"travel_booking/databases"
	"travel_booking/typ"
)

func CreateReservation(req typ.CreateReservationReq) error {
	return databases.CreateReservation(req)
}
