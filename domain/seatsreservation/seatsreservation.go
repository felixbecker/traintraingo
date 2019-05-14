package seatsreservation

import (
	"traintraingo/domain/ports"
	"traintraingo/domain/train"
)

type ticketmanager struct {
	trainDataService        ports.TrainDataService
	bookingReferenceService ports.BookingReferenceService
}

//New creates a new web ticket manager with dependencies to train data service and booking reference service
func New(trainDataService ports.TrainDataService, bookingReferenceService ports.BookingReferenceService) ports.SeatReserver {

	return &ticketmanager{
		trainDataService:        trainDataService,
		bookingReferenceService: bookingReferenceService,
	}
}

func (tm *ticketmanager) Reserve(trainID train.ID, numberOfSeats int) *train.Reservation {

	var failedReservation train.Reservation

	selectedTrain := tm.trainDataService.GetTrain(trainID)

	if selectedTrain.DoesNotExceedOveralTrainCapacity(numberOfSeats) {
		reservationAttempt, err := selectedTrain.BuildReservationAttempt(trainID, numberOfSeats)
		if err != nil {
			failedReservation := train.NewFailedReservation(trainID, train.EmptyBookingReference())
			return &failedReservation
		}
		if reservationAttempt.IsFullfilled() {

			bookingRef := tm.bookingReferenceService.GetBookingReference()
			reservationAttempt.AssignBookingReference(bookingRef)

			err := tm.trainDataService.BookSeats(trainID, bookingRef, reservationAttempt.Seats())
			if err == nil {
				reservation := reservationAttempt.Confirm()
				return &reservation
			}

		}

	}
	failedReservation = train.NewFailedReservation(trainID, train.EmptyBookingReference())
	return &failedReservation
}
