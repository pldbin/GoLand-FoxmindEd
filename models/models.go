package models

// Define a struct to hold information about a single holiday
type Holiday struct {
	Name        string `json:"name"`
	NameLocal   string `json:"name_local"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Country     string `json:"country"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	DateYear    string `json:"date_year"`
	DateMonth   string `json:"date_month"`
	DateDay     string `json:"date_day"`
	WeekDay     string `json:"week_day"`
}
