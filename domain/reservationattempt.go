package domain

//ReservationAttempt is a struct to represent a reservation attempts
type ReservationAttempt struct {
	numberOfRequestedSeats int
	seats                  []*Seat
	trainID                string
	bookingReference       string
}

//NewFailedReservationAttempt returns a Reservationattempt with no seats
func NewFailedReservationAttempt(trainID string, numberOfRequestedSeats int) ReservationAttempt {
	return ReservationAttempt{
		trainID:                trainID,
		numberOfRequestedSeats: numberOfRequestedSeats,
	}
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
func (r *ReservationAttempt) AssignBookingReference(bookingRef string) {
	for _, s := range r.seats {
		s.BookingRef = bookingRef
	}

	r.bookingReference = bookingRef
}

//Confirm confirms transforms a reservation attempt to a reservation
func (r *ReservationAttempt) Confirm() Reservation {
	return NewReservation(r.trainID, r.bookingReference, r.seats)
}
