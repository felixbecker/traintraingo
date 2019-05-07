package domain

type ReservationAttempt struct {
	numberOfRequestedSeats int
	seats                  []Seat
}

func NewFailedReservationAttempt(numberOfRequestedSeats int) ReservationAttempt {
	return ReservationAttempt{
		numberOfRequestedSeats: numberOfRequestedSeats,
	}
}
func (r *ReservationAttempt) Seats() []Seat {
	return r.seats
}

//IsFullfilled is true when the numberOfRequestedSeats matches the count of the seats in the reservation attempt
func (r *ReservationAttempt) IsFullfilled() bool {
	return len(r.seats) == r.numberOfRequestedSeats
}

func (r *ReservationAttempt) AssignBookingReference(bookingRef string) {
	for _, s := range r.seats {
		s.BookingRef = bookingRef
	}
}

func (r *ReservationAttempt) Confirm() Reservation {
	return NewReservation(r.seats)
}
