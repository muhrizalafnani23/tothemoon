package flight

import (
	"errors"
	"github.com/google/uuid"
)

type Flight struct {
	ID             string `json:"id"`
	FlightCode     string `json:"flight_code"`
	Departure      string `json:"departure"`
	Arrival        string `json:"arrival"`
	TotalPassenger int64  `json:"total_passenger"`
}

var listOfFLight = []Flight{
	{uuid.New().String(), "JT-256", "CGK", "DPS", 212},
	{uuid.New().String(), "GA-134", "CGK", "KNO", 154},
	{uuid.New().String(), "NAM-56", "DPS", "SIN", 130},
	{uuid.New().String(), "ID-123", "PLG", "DPS", 100},
}

func GetListFlight() []Flight {

	return listOfFLight
}

func GetFlightByCode(code string) (Flight, error) {
	for _, flight := range listOfFLight {
		if code == flight.FlightCode {
			return flight, nil
		}
	}

	return Flight{}, errors.New("flight not found")
}

func InsertNewFlight(flight Flight) ([]Flight, error) {
	for _, f := range listOfFLight {
		if flight.FlightCode == f.FlightCode {
			return nil, errors.New("duplicate code")
		}
	}

	flight.ID = uuid.New().String()
	listOfFLight = append(listOfFLight, flight)

	return listOfFLight, nil
}
