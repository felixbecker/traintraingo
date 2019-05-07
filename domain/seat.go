package domain

//Seat is a value type to represent seat in a traintrain domain
type Seat struct {
	coachName  string
	seatNumber int
	BookingRef string
}

//CoachName retrieves the coach name of a seat
func (s *Seat) CoachName() string {
	return s.coachName
}

//SeatNumber retrieves the seat number of a seat
func (s *Seat) SeatNumber() int {
	return s.seatNumber
}
