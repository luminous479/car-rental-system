package rental

type Rental struct {
	ID         int     `json:"id"`
	CarID      int     `json:"car_id"`
	Customer   string  `json:"customer"`
	Days       int     `json:"days"`
	TotalPrice float64 `json:"total_price"`
	Returned   bool    `json:"returned"`
}