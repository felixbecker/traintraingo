package infra

import (
	"encoding/json"
	"traintraingo/domain"
)

type ReservationRestAdapter interface {
}

type BookingReferenceAdapter interface {
}

type TrainDataServiceAdapter interface {
}

type BookingServiceAdapter interface {
}

type SeatReservationAdapter interface {
	Post(dto ReservationRequestDto) json.RawMessage
}

//ReservationRequestDto a Dto to request a reservation
type ReservationRequestDto struct {
	trainID       string
	numberOfSeats int
}

type seatReservationAdapter struct {
	hexagon domain.SeatReserver
}

func (s *seatReservationAdapter) Post(dto ReservationRequestDto) json.RawMessage {
	// adapt from Infra to Domain
	numberOfSeatsToRequest := dto.numberOfSeats
	trainID := dto.trainID

	// Call business logic
	json := s.hexagon.Reserve(trainID, numberOfSeatsToRequest)
	// Adapt from Domain to Infra

}

func NewSeatReservationAdapter(hexagon domain.SeatReserver) SeatReservationAdapter {

}
