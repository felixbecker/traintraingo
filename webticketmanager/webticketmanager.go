package webticketmanager

import "traintraingo/domain"

type ticketmanager struct {
	trainDataService        TrainDataService
	bookingReferenceService BookingReferenceService
}

//New creates a new web ticket manager with dependencies to train data service and booking reference service
func New(trainDataService TrainDataService, bookingReferenceService BookingReferenceService) domain.SeatReserver {

	return &ticketmanager{
		trainDataService:        trainDataService,
		bookingReferenceService: bookingReferenceService,
	}
}

func (tm *ticketmanager) Reserve(trainID string, numberOfSeats int) domain.Reservation {

	train := tm.trainDataService.Train(trainID)
	reservationAttempt := train.BuildReservationAttempt(numberOfSeats)

	if reservationAttempt.IsFullfilled() {
		bookingRef := tm.bookingReferenceService.GetBookingReference()
		err := tm.trainDataService.BookSeats(trainID, bookingRef, reservationAttempt.Seats())
		if err != nil {
			return reservationAttempt.Confirm()
		}
	}
	return ""
}
