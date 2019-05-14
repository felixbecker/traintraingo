package traindataservice

import (
	"traintraingo/domain/ports"
	"traintraingo/domain/train"
)

type dataservice struct {
}

//New creates a new train data service
func New() ports.TrainDataService {
	return &dataservice{}
}

func (ds *dataservice) GetTrain(id train.ID) train.Train {
	return train.Train{}
}
func (ds *dataservice) AdaptTrainTopology(jsonTopology string) []train.Seat {

	return []train.Seat{}
}

func (ds *dataservice) BookSeats(trainID train.ID, bookingReference train.BookingReference, seats []*train.Seat) error {
	return nil
}
