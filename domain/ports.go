package domain

//TrainID is a type for the train id
type TrainID string

//SeatReserver is an interface to abstract the web ticket manager type
type SeatReserver interface {
	Reserve(id TrainID, numberOfSeats int) Reservation
}

//TrainDataService is a port
type TrainDataService interface {
	GetTrain(id TrainID) Train
	BookSeats(trainID TrainID, bookingReference string, seats []*Seat) error
}

//BookingReferenceService is a port
type BookingReferenceService interface {
	GetBookingReference() string
}
