package seatsreservation

import (
	"traintraingo/domain"
)

type ticketmanager struct {
	trainDataService        domain.TrainDataService
	bookingReferenceService domain.BookingReferenceService
}

//New creates a new web ticket manager with dependencies to train data service and booking reference service
func New(trainDataService domain.TrainDataService, bookingReferenceService domain.BookingReferenceService) domain.SeatReserver {

	return &ticketmanager{
		trainDataService:        trainDataService,
		bookingReferenceService: bookingReferenceService,
	}
}

func (tm *ticketmanager) Reserve(trainID domain.TrainID, numberOfSeats int) domain.Reservation {

	train := tm.trainDataService.GetTrain(trainID)

	if train.DoesNotExceedOveralTrainCapacity(numberOfSeats) {
		reservationAttempt := train.BuildReservationAttempt(trainID, numberOfSeats)
		if reservationAttempt.IsFullfilled() {

			bookingRef := tm.bookingReferenceService.GetBookingReference()
			reservationAttempt.AssignBookingReference(bookingRef)

			err := tm.trainDataService.BookSeats(trainID, bookingRef, reservationAttempt.Seats())
			if err == nil {
				reservation := reservationAttempt.Confirm()
				return reservation
			}
		}
	}

	return domain.NewFailedReservation(trainID)
}
