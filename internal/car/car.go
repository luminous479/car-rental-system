package car

type Car struct {
	ID        int     `json:"id"`
	Brand     string  `json:"brand"`
	Model     string  `json:"model"`
	Year      int     `json:"year"`
	DailyRate float64 `json:"daily_rate"`
	Available bool    `json:"available"`
}