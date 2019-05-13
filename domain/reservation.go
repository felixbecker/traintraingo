package domain

//Reservation struct to hold information for a reservation
type Reservation struct {
	trainID          TrainID
	bookingReference BookingReference
	seats            []*Seat
}

//TrainID imutable getter for the train id
func (r *Reservation) TrainID() TrainID {
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
func NewReservation(trainID TrainID, bookingReference BookingReference, seats []*Seat) Reservation {
	return Reservation{
		trainID:          trainID,
		bookingReference: bookingReference,
		seats:            seats,
	}
}

//NewFailedReservation returns an empty reservation
func NewFailedReservation(trainID TrainID) Reservation {
	return Reservation{trainID: trainID}
}
