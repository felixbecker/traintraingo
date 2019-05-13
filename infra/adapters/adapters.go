package adapters

import (
	"encoding/json"
	"fmt"
	"strconv"
	"traintraingo/domain"
)

type seatDto struct {
	CoachName  string `json:"coach"`
	SeatNumber string `json:"seat_number"`
	BookingRef string `json:"booking_reference"`
}

type trainDto struct {
	Seats map[string]seatDto `json:"seats"`
}

//AdaptTrainTopology adapts a json string into a list of seats
func AdaptTrainTopology(jsonString string) ([]*domain.Seat, error) {

	seats := []*domain.Seat{}
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
		seats = append(seats, domain.NewSeat(seat.CoachName, seatNumber, seat.BookingRef))
	}
	return seats, nil
}

type reservationDto struct {
	TrainID          string   `json:"train_id"`
	BookingReference string   `json:"booking_reference"`
	Seats            []string `json:"seats"`
}

//AdaptTrainID adapts the train id and returns its string representation
func AdaptTrainID(trainID domain.TrainID) string {
	return string(trainID)
}

//AdaptReservation adapts the reservation and returns its json bytes representation
func AdaptReservation(reservation domain.Reservation) []byte {

	dto := reservationDto{}
	dto.TrainID = AdaptTrainID(reservation.TrainID())
	dto.BookingReference = reservation.BookingReference()
	seats := []string{}
	for _, seat := range reservation.Seats() {
		seats = append(seats, fmt.Sprintf("%d%s", seat.SeatNumber(), seat.CoachName()))
	}
	dto.Seats = seats
	jsonBytes, err := json.Marshal(dto)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return jsonBytes
}
