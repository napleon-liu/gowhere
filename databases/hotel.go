package databases

import "travel_booking/typ"

func CreateHotel(req typ.CreateHotelReq) error {
	_, err := Client.Exec("INSERT INTO hotels (location, price, numRooms, numAvail) VALUES (?,?,?,?)", req.Location, req.Price, req.NumRooms, req.NumAvail)
	return err
}

func DeleteHotel(req typ.DeleteHotelReq) error {
	_, err := Client.Exec("DELETE FROM hotels WHERE location=?", req.Location)
	return err
}

func UpdateHotel(req typ.UpdateHotelReq) error {
	return nil
}

func GetHotel() (typ.GetHotelRes, error) {
	var res typ.GetHotelRes
	rows, err := Client.Query("SELECT * FROM hotels")
	for rows.Next() {
		var hotel typ.Hotel
		err = rows.Scan(&hotel.Location, &hotel.Price, &hotel.NumRooms, &hotel.NumAvail)
		if err != nil {
			return res, err
		}
		res.Hotels = append(res.Hotels, hotel)
	}
	return res, err
}
