package train

//Reservation struct to hold information for a reservation
type Reservation struct {
	trainID          ID
	bookingReference BookingReference
	seats            []*Seat
}

//TrainID imutable getter for the train id
func (r *Reservation) TrainID() ID {
	return r.trainID
}

//BookingReference imutable getter for the booking reference
func (r *Reservation) BookingReference() BookingReference {
	return r.bookingReference
}

//Seats imutable getter for the seats
func (r *Reservation) Seats() []*Seat {
	return r.seats
}

//NewReservation creates a new reservation based on the trainID, bookingReference and the seats
func newReservation(trainID ID, bookingReference BookingReference, seats []*Seat) Reservation {
	return Reservation{
		trainID:          trainID,
		bookingReference: bookingReference,
		seats:            seats,
	}
}

//NewFailedReservation creates a reservation that holds only trainID and booking reference since it failed
func newFailedReservation(trainID ID, bookingReference BookingReference) Reservation {
	return Reservation{
		trainID:          trainID,
		bookingReference: bookingReference,
	}
}
