package databases

import "travel_booking/typ"

func CreateReservation(req typ.CreateReservationReq) error {
	Client.Begin()
	defer Client.Close()
	Client.Exec("INSERT INTO reservation (custName, resvType, resvKey) VALUES (?, ?, ?)", req.CustName, req.ResvType, req.ResvKey)
	switch req.ResvType {
	case 1: // flight
		_, err := Client.Exec("UPDATE flight SET numAvail = numAvail - 1 WHERE flightNum = ?", req.ResvKey)
		return err
	case 2: // hotel
		_, err := Client.Exec("UPDATE hotel SET numAvail = numAvail - 1 WHERE location = ?", req.ResvKey)
		return err
	case 3: // bus
		_, err := Client.Exec("UPDATE bus SET numAvail = numAvail - 1 WHERE location = ?", req.ResvKey)
		return err
	}
	return nil
}
