package services

import (
	"travel_booking/databases"
	"travel_booking/typ"
)

func CreateBus(req typ.CreateBusReq) error {
	return databases.CreateBus(req)
}

func DeleteBus(req typ.DeleteBusReq) error {
	return databases.DeleteBus(req)
}

func UpdateBus(req typ.UpdateBusReq) error {
	return databases.UpdateBus(req)
}

func GetBus() (typ.GetBusRes, error) {
	return databases.GetBus()
}
