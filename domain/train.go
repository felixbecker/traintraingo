package domain

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

//NewTrain creates a new train based on a set of seats
func NewTrain(seats []Seat) Train {
	return Train{
		seats: seats,
	}
}
