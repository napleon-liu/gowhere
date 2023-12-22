package services

import (
	"travel_booking/databases"
	"travel_booking/typ"
)

func CreateHotel(req typ.CreateHotelReq) error {
	return databases.CreateHotel(req)
}

func DeleteHotel(req typ.DeleteHotelReq) error {
	return databases.DeleteHotel(req)
}

func UpdateHotel(req typ.UpdateHotelReq) error {
	return databases.UpdateHotel(req)
}

func GetHotel() (typ.GetHotelRes, error) {
	return databases.GetHotel()
}
