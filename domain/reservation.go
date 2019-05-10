package domain

//Reservation struct to hold information for a reservation
type Reservation struct {
	trainID          string
	bookingReference string
	seats            []*Seat
}

func (r *Reservation) TrainID() string {
	return r.trainID
}
func (r *Reservation) BookingReference() string {
	return r.bookingReference
}

func (r *Reservation) Seats() []*Seat {
	return r.seats
}

//NewReservation creates a new reservation based on the trainID, bookingReference and the seats
func NewReservation(trainID string, bookingReference string, seats []*Seat) Reservation {
	return Reservation{
		trainID:          trainID,
		bookingReference: bookingReference,
		seats:            seats,
	}
}

//NewFailedReservation returns an empty reservation
func NewFailedReservation(trainID string) Reservation {
	return Reservation{trainID: trainID}
}
