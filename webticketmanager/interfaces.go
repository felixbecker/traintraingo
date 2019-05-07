package webticketmanager

//TrainDataService is a port
type TrainDataService interface {
	Train(id string) []domain.Seat
}

//BookingReferenceService is a port
type BookingReferenceService interface {
}
