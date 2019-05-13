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

func (ds *dataservice) GetTrain(id domain.TrainID) domain.Train {
	return domain.Train{}
}
func (ds *dataservice) AdaptTrainTopology(jsonTopology string) []domain.Seat {

	return []domain.Seat{}
}

func (ds *dataservice) BookSeats(trainID domain.TrainID, bookingReference domain.BookingReference, seats []*domain.Seat) error {
	return nil
}
