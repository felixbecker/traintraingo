package domain

//SeatReserver is an interface to abstract the web ticket manager type
type SeatReserver interface {
	Reserve(trainID string, numberOfSeats int) Reservation
}
