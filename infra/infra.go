package infra

import (
	"encoding/json"
	"traintraingo/domain/ports"
	"traintraingo/infra/adapters"
)

//ReservationRequestDto a Dto to request a reservation
type ReservationRequestDto struct {
	TrainID       string
	NumberOfSeats int
}

//SeatReservationAdapter is an adapter to abstract the domain from the infrastructure
type SeatReservationAdapter interface {
	Post(dtp ReservationRequestDto) (json.RawMessage, error)
}
type seatReservationAdapter struct {
	hexagon ports.SeatReserver
}

func (s *seatReservationAdapter) Post(dto ReservationRequestDto) (json.RawMessage, error) {
	// adapt from Infra to Domain
	numberOfSeatsToRequest := dto.NumberOfSeats
	trainID := adapters.AdaptTrainIDString(dto.TrainID)

	// Call business logic
	reservation := s.hexagon.Reserve(trainID, numberOfSeatsToRequest)

	// Adapt from Domain to Infra

	return adapters.AdaptReservation(*reservation)

}

//NewSeatReservationAdapter constructs a new seat reservation adapter
func NewSeatReservationAdapter(hexagon ports.SeatReserver) SeatReservationAdapter {
	return &seatReservationAdapter{
		hexagon: hexagon,
	}
}
