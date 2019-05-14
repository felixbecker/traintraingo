package train

import "sort"

//NewCoach creates a Coach struct
func NewCoach(coachName string) *Coach {
	return &Coach{
		name: coachName,
	}
}

//Coach ist a struct holder information about a coach
type Coach struct {
	name  string
	seats []*Seat
}

//AddSeat allows to add a seat to a coach
func (c *Coach) AddSeat(seat *Seat) {
	c.seats = append(c.seats, seat)
}

//Seats get all seats for a coach
func (c *Coach) Seats() []*Seat {
	return c.seats
}

//BuildReservationAttempt creates a reservation attempt based on the number of requested seats
func (c *Coach) BuildReservationAttempt(trainID ID, numberOfRequestedSeats int) ReservationAttempt {

	avaibleSeats := []*Seat{}
	sort.Slice(c.seats, func(i, j int) bool {
		return c.seats[i].seatNumber < c.seats[j].seatNumber
	})
	for idx, seat := range c.seats {
		if seat.IsAvailable() {
			if idx+1 <= numberOfRequestedSeats {
				avaibleSeats = append(avaibleSeats, seat)
			}
		}

	}

	return ReservationAttempt{
		trainID:                trainID,
		numberOfRequestedSeats: numberOfRequestedSeats,
		seats:                  avaibleSeats,
	}

}
