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

	train := tm.trainDataService.Train(trainID)

	reservationAttempt := train.BuildReservationAttempt(numberOfSeats)

	if reservationAttempt.IsFullfilled() {
		bookingRef := tm.bookingReferenceService.GetBookingReference()

		tm.trainDataService.BookSeats(trainID, bookingRef, reservationAttempt.Seats())

	}
	return ""
}
