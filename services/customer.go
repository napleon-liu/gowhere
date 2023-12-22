package services

import (
	"travel_booking/databases"
	"travel_booking/typ"
)

func CreateCustomer(req typ.CreateCustomerReq) error {
	return databases.CreateCustomer(req)
}

func DeleteCustomer(req typ.DeleteCustomerReq) error {
	return databases.DeleteCustomer(req.CustName)
}

func UpdateCustomer(req typ.UpdateCustomerReq) error {
	return databases.UpdateCustomer(req)
}

func GetCustomer() (typ.GetCustomerRes, error) {
	return databases.GetCustomer()
}
