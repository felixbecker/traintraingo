package adapters

import (
	"encoding/json"
	"fmt"
	"strconv"
	"traintraingo/domain"
)

type SeatDto struct {
	CoachName  string `json:"coach"`
	SeatNumber string `json:"seat_number"`
}

type trainDto struct {
	Seats map[string]SeatDto `json:"seats"`
}

func AdaptTrainTopology(jsonString string) ([]domain.Seat, error) {

	seats := []domain.Seat{}
	var train trainDto
	//err := json.NewDecoder(jsonResponse).Decode(&train)
	err := json.Unmarshal([]byte(jsonString), &train)
	if err != nil {
		return nil, err
	}

	for _, seat := range train.Seats {
		seatNumber, err := strconv.Atoi(seat.SeatNumber)
		if err != nil {
			return nil, fmt.Errorf("malicious seat number")
		}
		seats = append(seats, domain.NewSeat(seat.CoachName, seatNumber))
	}
	return seats, nil
}
