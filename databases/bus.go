package databases

import "travel_booking/typ"

func GetBus() (typ.GetBusRes, error) {
	var res typ.GetBusRes
	rows, err := Client.Query("SELECT * FROM bus")
	for rows.Next() {
		var bus typ.Bus
		err = rows.Scan(&bus.Location, &bus.Price, &bus.NumSeats, &bus.NumAvail)
		res.Buses = append(res.Buses, bus)
	}
	return res, err
}

func CreateBus(req typ.CreateBusReq) error {
	_, err := Client.Exec("INSERT INTO bus (location, price, numBus, numAvail) VALUES (?,?,?,?)", req.Location, req.Price, req.NumBus, req.NumAvail)
	return err
}

func DeleteBus(req typ.DeleteBusReq) error {
	_, err := Client.Exec("DELETE FROM bus WHERE location=?", req.Location)
	return err
}

func UpdateBus(req typ.UpdateBusReq) error {
	if req.Price != 0 {
		Client.Exec("UPDATE bus SET price=? where location=?", req.Price, req.Location)
	}
	if req.NumAvail != 0 {
		Client.Exec("UPDATE bus SET numAvail=? where location=?", req.NumAvail, req.Location)
	}
	if req.NumBus != 0 {
		Client.Exec("UPDATE bus SET numBus=? where location=?", req.NumBus, req.Location)
	}
	return nil
}
