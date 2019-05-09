package domain

//Reservation struct to hold information for a reservation
type Reservation struct {
	trainID          string
	bookingReference string
	seats            []Seat
}

//NewReservation creates a new reservation based on the trainID, bookingReference and the seats
func NewReservation(trainID string, bookingReference string, seats []Seat) Reservation {
	return Reservation{
		trainID:          trainID,
		bookingReference: bookingReference,
		seats:            seats,
	}
}

//NewFailedReservation returns an empty reservation
func NewFailedReservation() Reservation {
	return Reservation{}
}
