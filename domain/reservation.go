package domain

type Reservation struct {
	trainID          string
	bookingReference string
	seats            []Seat
}

func NewReservation(trainID string, bookingReference string, seats []Seat) Reservation {
	return Reservation{
		trainID:          trainID,
		bookingReference: bookingReference,
		seats:            seats,
	}
}
