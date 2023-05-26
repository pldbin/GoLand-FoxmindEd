package holiday

import (
	"About_me_Bot/models"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

// Define a struct to hold the response from the Holidays API
type HolidaysResponse []models.Holiday

// Before you use fmt.Sprintf(), write valildator function. Values for current variables can be empty....
func validateArgs(apiKey, country string, year, month, day int) error {
	if apiKey == "" {
		return fmt.Errorf("API key is required")
	}

	if country == "" {
		return fmt.Errorf("Country is required")
	}

	if year < 0 {
		return fmt.Errorf("Year must be a positive integer")
	}

	if month < 1 || month > 12 {
		return fmt.Errorf("Month must be between 1 and 12")
	}

	if day < 1 || day > 31 {
		return fmt.Errorf("Day must be between 1 and 31")
	}

	if err := validateArgs(apiKey, country, year, month, day); err != nil {
		log.Fatalf("Invalid arguments: %v", err)
	}

	return nil
}

func GetHolidays(countryCode string, date time.Time, apiKey string) ([]models.Holiday, error) {

	// Build the URL for the Holidays API endpoint
	url := fmt.Sprintf("https://holidays.abstractapi.com/v1/?api_key=%s&country=%s&year=%d&month=%d&day=%d",
		apiKey,
		countryCode,
		date.Year(),
		int(date.Month()),
		date.Day(),
	)

	// Send an HTTP GET request to the Holidays API endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decode the JSON response from the Holidays API
	var holidays []models.Holiday
	err = json.Unmarshal(body, &holidays)
	if err != nil {
		return nil, err
	}

	// Return the list of holidays for the specified country and date
	return holidays, nil
}
