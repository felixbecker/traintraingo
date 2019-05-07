package traindataservice

import (
	"traintraingo/domain"
	"traintraingo/webticketmanager"
)

type dataservice struct {
}

func New() webticketmanager.TrainDataService {
	return &dataservice{}
}

type SeatJsonPoco struct {
}

func (ds *dataservice) Train(id string) domain.Train {
	return domain.Train{}
}
func (ds *dataservice) AdaptTrainTopology(jsonTopology string) []domain.Seat {

	return []domain.Seat{}
}

func (ds *dataservice) BookSeats(trainID string, bookingReference string, seats []domain.Seat) {

}
