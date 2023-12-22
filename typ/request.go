package typ

type CreateCustomerReq struct {
	CustName string `json:"custName"`
	CustID   string `json:"custNid"`
}

type DeleteCustomerReq struct {
	CustName string `json:"custName"`
}

type UpdateCustomerReq struct {
	NewName string `json:"newName"`
	OldName string `json:"oldName"`
}

type GetCustomerReq struct {
}

type CreateBusReq struct {
	Location string `json:"location"`
	Price    string `json:"price"`
	NumBus   string `json:"numBus"`
	NumAvail string `json:"numAvail"`
}

type DeleteBusReq struct {
	Location string `json:"location"`
}

type UpdateBusReq struct {
	Location string `json:"location"`
	Price    int    `json:"price"`
	NumBus   int    `json:"numBus"`
	NumAvail int    `json:"numAvail"`
}

type GetBusReq struct {
}

type CreateFlightReq struct {
	FlightNum string `json:"flightNum"`
	Price     int    `json:"price"`
	NumSeats  int    `json:"numSeats"`
	NumAvail  int    `json:"numAvail"`
	FromCity  string `json:"fromCity"`
	ArivCity  string `json:"arivCity"`
}

type DeleteFlightReq struct {
	FlightNum string `json:"flightNum"`
}

type UpdateFlightReq struct {
}

type GetFlightReq struct {
}

type CreateHotelReq struct {
	Location string `json:"location"`
	Price    int    `json:"price"`
	NumRooms int    `json:"numRooms"`
	NumAvail int    `json:"numAvail"`
}

type DeleteHotelReq struct {
	Location string `json:"location"`
}

type UpdateHotelReq struct {
}

type GetHotelReq struct {
}

type CreateReservationReq struct {
	CustName string `json:"custName"`
	ResvType int    `json:"resvType"`
	ResvKey  string `json:"resvKey"`
}

type DeleteReservationReq struct {
}

type UpdateReservationReq struct {
}

type GetReservationReq struct {
}
