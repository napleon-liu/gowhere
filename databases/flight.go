package databases

import "travel_booking/typ"

func CreateFlight(req *typ.CreateFlightReq) error {
	_, err := Client.Exec("INSERT INTO flight VALUES (?, ?, ?, ?, ?, ?)", req.FlightNum, req.Price, req.NumSeats, req.NumAvail, req.FromCity, req.ArivCity)
	return err
}

func DeleteFlight(req *typ.DeleteFlightReq) error {
	_, err := Client.Exec("DELETE FROM flight WHERE flightNum = ?", req.FlightNum)
	return err
}

func UpdateFlight(req *typ.UpdateFlightReq) error {
	return nil
}

func GetFlight() (typ.GetFlightRes, error) {
	var res typ.GetFlightRes
	result, err := Client.Query("SELECT * FROM flights")
	for result.Next() {
		var flight typ.Flight
		result.Scan(&flight.FlightNum, &flight.Price, &flight.NumSeats, &flight.NumAvail, &flight.FromCity, &flight.ArivCity)
		res.Flights = append(res.Flights, flight)
	}
	return res, err
}
