package typ

// MetaData 响应元数据
type MetaData struct {
	Code int    `json:"HTTPCode"`
	Msg  string `json:"Msg"`
}

// T HTTP响应，包含MetaData元数据，及Result返回结果。
type T struct {
	MetaData *MetaData   `json:"ResponseMetadata"`
	Result   interface{} `json:"Result"`
}

// Flight type flight FLIGHTS (String flightNum, int price, int numSeats, int numAvail, String FromCity, String ArivCity)；
type Flight struct {
	FlightNum string `json:"flightNum"`
	Price     int    `json:"price"`
	NumSeats  int    `json:"numSeats"`
	NumAvail  int    `json:"numAvail"`
	FromCity  string `json:"fromCity"`
	ArivCity  string `json:"arivCity"`
}

type GetFlightRes struct {
	Flights []Flight `json:"flights"`
}

// Hotel type hotel HOTELS (String location, int price, int numRooms, int numAvail)；
type Hotel struct {
	Location string `json:"location"`
	Price    int    `json:"price"`
	NumRooms int    `json:"numRooms"`
	NumAvail int    `json:"numAvail"`
}

type GetHotelRes struct {
	Hotels []Hotel `json:"hotels"`
}

type Bus struct {
	Location string `json:"location"`
	Price    int    `json:"price"`
	NumSeats int    `json:"numSeats"`
	NumAvail int    `json:"numAvail"`
}

type GetBusRes struct {
	Buses []Bus `json:"buses"`
}

// Customer CUSTOMER (String custName, String custID)；
type Customer struct {
	CustName string `json:"custName"`
	CustID   string `json:"custID"`
}

type GetCustomerRes struct {
	Customers []Customer `json:"customers"`
}
