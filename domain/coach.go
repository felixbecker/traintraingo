package domain

func NewCoach(coachName string) Coach {
	return Coach{
		name: coachName,
	}
}

type Coach struct {
	name  string
	seats []Seat
}

func (c Coach) AddSeat(seat Seat) {
	c.seats = append(c.seats, seat)
}

func (c Coach) Seats() []Seat {
	return c.seats
}

func (c Coach) BuildReservationAttempt(numberOfRequestedSeats int) ReservationAttempt {
	avaibleSeats := []Seat{}
	for idx, seat := range c.seats {
		if seat.IsAvailable() {
			if idx <= numberOfRequestedSeats {
				avaibleSeats = append(avaibleSeats, seat)
			}
		}

	}

	return ReservationAttempt{
		numberOfRequestedSeats: numberOfRequestedSeats,
		seats:                  avaibleSeats,
	}

}
