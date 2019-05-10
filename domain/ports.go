package domain

//SeatReserver is an interface to abstract the web ticket manager type
type SeatReserver interface {
	Reserve(trainID string, numberOfSeats int) Reservation
}

//TrainDataService is a port
type TrainDataService interface {
	GetTrain(id string) Train
	BookSeats(trainID string, bookingReference string, seats []*Seat) error
}

//BookingReferenceService is a port
type BookingReferenceService interface {
	GetBookingReference() string
}
