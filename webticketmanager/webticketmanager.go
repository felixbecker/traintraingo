package webticketmanager

//ReservationRequestDto a Dto to request a reservation
type ReservationRequestDto struct {
	trainID       string
	numberOfSeats int
}

//TicketManager is an interface to abstract the web ticket manager type
type TicketManager interface {
	Reserve(trainID string, numberOfSeats int) string
}

type ticketmanager struct {
	trainDataService        TrainDataService
	bookingReferenceService BookingReferenceService
}

//New creates a new web ticket manager with dependencies to train data service and booking reference service
func New(trainDataService TrainDataService, bookingReferenceService BookingReferenceService) TicketManager {

	return &ticketmanager{
		trainDataService:        trainDataService,
		bookingReferenceService: bookingReferenceService,
	}
}

func (tm *ticketmanager) Reserve(trainID string, numberOfSeats int) string {

	jsonTrain := tm.trainDataService.Train(trainID)
	train := domain.NewTrain(jsonTrain)
	_ = train
	return ""
}
