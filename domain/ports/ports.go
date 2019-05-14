package ports

import "traintraingo/domain/train"

//SeatReserver is an interface to abstract the web ticket manager type
type SeatReserver interface {
	Reserve(id train.ID, numberOfSeats int) *train.Reservation
}

//TrainDataService is a port
type TrainDataService interface {
	GetTrain(id train.ID) train.Train
	BookSeats(trainID train.ID, ref train.BookingReference, seats []*train.Seat) error
}

//BookingReferenceService is a port
type BookingReferenceService interface {
	GetBookingReference() train.BookingReference
}
