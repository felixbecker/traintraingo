package domain

//Reservationattempt holds information regarding a possible reservation
type ReservationAttempt struct {
	numberOfRequestedSeats int
	seats                  []Seat
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

//Train is value type thats hold train information
type Train struct {
	seats []Seat
}

//Seats retrieves all seats for a given train
func (t *Train) Seats() []Seat {
	return t.seats
}

//ReservedSeats retrieves all reserved seats
func (t *Train) ReservedSeats() []Seat {
	reservedSeats := []Seat{}
	for _, s := range t.seats {
		if len(s.BookingRef) > 0 {
			reservedSeats = append(reservedSeats, s)
		}
	}
	return reservedSeats
}

//MaxSeats the number of maximal seats for a given train
func (t *Train) MaxSeats() int {
	return len(t.Seats())
}

//DoesNotExceedOveralTrainCapacity checks whether the numberOfRequestedSeats exceeds the overall train capacity
func (t *Train) DoesNotExceedOveralTrainCapacity(numberOfRequestedSeats int) bool {
	return float64(len(t.ReservedSeats())+numberOfRequestedSeats) < float64(t.MaxSeats())*float64(0.70)
}
func (t *Train) BuildReservationAttempt(seatRequested int) ReservationAttempt {
	avaibleSeats := []Seat{}
	for idx, s := range t.seats {
		if len(s.BookingRef) == 0 {
			if idx <= seatRequested {
				avaibleSeats = append(avaibleSeats, s)
			}
		}

	}
	return ReservationAttempt{
		numberOfRequestedSeats: seatRequested,
		seats:                  avaibleSeats,
	}
}

//NewTrain creates a new train based on a set of seats
func NewTrain(seats []Seat) Train {
	return Train{
		seats: seats,
	}
}
