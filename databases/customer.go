package databases

import (
	"travel_booking/typ"
)

func CreateCustomer(req typ.CreateCustomerReq) error {
	_, err := Client.Exec("INSERT INTO customers (custName, custID) VALUES (?,?)", req.CustName, req.CustID)
	return err
}

func DeleteCustomer(custName string) error {
	_, err := Client.Exec("DELETE FROM customers WHERE custName=?", custName)
	return err
}

func UpdateCustomer(req typ.UpdateCustomerReq) error {
	_, err := Client.Exec("UPDATE customers SET custName=? where custNmae=?", req.NewName, req.OldName)
	return err
}

func GetCustomer() (typ.GetCustomerRes, error) {
	var res typ.GetCustomerRes
	result, err := Client.Query("SELECT * FROM customers")
	for result.Next() {
		var customer typ.Customer
		err = result.Scan(&customer.CustName, &customer.CustID)
		if err != nil {
			return res, err
		}
		res.Customers = append(res.Customers, customer)
	}
	return res, err
}
