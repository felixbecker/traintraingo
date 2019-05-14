package train

//Seat is a value type to represent seat in a traintrain domain
type Seat struct {
	coachName  string
	seatNumber int
	BookingRef BookingReference
}

//CoachName retrieves the coach name of a seat
func (s *Seat) CoachName() string {
	return s.coachName
}

//SeatNumber retrieves the seat number of a seat
func (s *Seat) SeatNumber() int {
	return s.seatNumber
}

//IsAvailable checks the availability for a seat which is determend by checking if the BookingRef is empty
func (s *Seat) IsAvailable() bool {
	return len(s.BookingRef) == 0
}

//NewSeat creates a new seat struct based on coachName and seatNumber
func NewSeat(coachName string, seatNumber int, bookingRef BookingReference) *Seat {
	return &Seat{
		coachName:  coachName,
		seatNumber: seatNumber,
		BookingRef: bookingRef,
	}
}
