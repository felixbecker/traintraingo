package train

import (
	"fmt"
	"math"
)

//ID is a type for the train id
type ID string

//BookingReference is atype for the booking reference
type BookingReference string

//Train is value type thats hold train information
type Train struct {
	Coaches map[string]*Coach
}

//EmptyBookingReference returns an empty booking reference
func EmptyBookingReference() BookingReference {
	return BookingReference("")
}

//Seats retrieves all seats for a given train
func (t *Train) Seats() []*Seat {
	seats := []*Seat{}
	for _, coach := range t.Coaches {
		seats = append(seats, coach.seats...)
	}
	return seats
}

//ReservedSeats retrieves all reserved seats
func (t *Train) ReservedSeats() []*Seat {
	reservedSeats := []*Seat{}
	for _, s := range t.Seats() {
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
	result := float64(len(t.ReservedSeats())+numberOfRequestedSeats) <= math.Floor(float64(t.MaxSeats())*float64(0.70))
	fmt.Println(t.ReservedSeats())
	return result
}

//BuildReservationAttempt creates a reservation attempt from the number of requested seats
func (t *Train) BuildReservationAttempt(trainID ID, seatRequested int) (*ReservationAttempt, error) {

	for _, coach := range t.Coaches {
		reservationAttempt := coach.BuildReservationAttempt(trainID, seatRequested)
		if reservationAttempt.IsFullfilled() {

			return &reservationAttempt, nil
		}
	}
	return nil, fmt.Errorf("reservation attempt to reserve %d seat(s) failed for train: %s", seatRequested, trainID)
}

//New creates a new train based on a set of seats
func New(seats []*Seat) Train {
	coaches := map[string]*Coach{}
	for _, seat := range seats {

		if _, ok := coaches[seat.CoachName()]; !ok {
			coaches[seat.CoachName()] = NewCoach(seat.CoachName())
		}
		coaches[seat.CoachName()].AddSeat(seat)

	}

	return Train{
		Coaches: coaches,
	}
}
