package train

//ReservationAttempt is a struct to represent a reservation attempts
type ReservationAttempt struct {
	numberOfRequestedSeats int
	seats                  []*Seat
	trainID                ID
	bookingReference       BookingReference
}

//Seats returns all seats for a reservation attempts
func (r *ReservationAttempt) Seats() []*Seat {
	return r.seats
}

//IsFullfilled is true when the numberOfRequestedSeats matches the count of the seats in the reservation attempt
func (r *ReservationAttempt) IsFullfilled() bool {
	return len(r.seats) == r.numberOfRequestedSeats
}

//AssignBookingReference assigns the booking reference for each seat in the reservation attempt
func (r *ReservationAttempt) AssignBookingReference(bookingRef BookingReference) {
	for _, s := range r.seats {
		s.BookingRef = bookingRef
	}

	r.bookingReference = bookingRef
}

//Confirm confirms transforms a reservation attempt to a reservation
func (r *ReservationAttempt) Confirm() Reservation {
	return newReservation(r.trainID, r.bookingReference, r.seats)
}
