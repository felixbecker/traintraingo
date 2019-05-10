package traindataservice

import (
	"traintraingo/domain"
)

type dataservice struct {
}

//New creates a new train data service
func New() domain.TrainDataService {
	return &dataservice{}
}

//SeatJSONPoco value object to hold information for serialization
type SeatJSONPoco struct {
}

func (ds *dataservice) Train(id string) domain.Train {
	return domain.Train{}
}
func (ds *dataservice) AdaptTrainTopology(jsonTopology string) []domain.Seat {

	return []domain.Seat{}
}

func (ds *dataservice) BookSeats(trainID string, bookingReference string, seats []*domain.Seat) error {
	return nil
}
