package webticketmanager

import "traintraingo/domain"

//TrainDataService is a port
type TrainDataService interface {
	Train(id string) domain.Train
	AdaptTrainTopology(jsonTopology string) []domain.Seat
	BookSeats(trainID string, bookingReference string, seats []domain.Seat)
}

//BookingReferenceService is a port
type BookingReferenceService interface {
	GetBookingReference() string
}
