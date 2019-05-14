package train_test

import (
	"testing"
	"traintraingo/domain/train"
	"traintraingo/infra/adapters"
)

const TrainIDconst string = "9043-2017-09-22"
const BookingReference string = "75bcd15"

func Test_Build_Reservation_Attempt(t *testing.T) {

	seats := []*train.Seat{
		train.NewSeat("A", 1, ""),
		train.NewSeat("A", 2, ""),
		train.NewSeat("A", 3, ""),
		train.NewSeat("A", 4, ""),
		train.NewSeat("A", 5, ""),
		train.NewSeat("A", 6, ""),
		train.NewSeat("A", 7, ""),
		train.NewSeat("A", 8, ""),
		train.NewSeat("A", 9, ""),
		train.NewSeat("A", 10, ""),
	}

	train := train.New(seats)

	expectedNumberOfSeatsForCoachA := 10

	numberOfSeats := len(train.Coaches["A"].Seats())
	if numberOfSeats != expectedNumberOfSeatsForCoachA {
		t.Errorf("Epected the number of seats in coach 'A' to be %d; got: %d",
			expectedNumberOfSeatsForCoachA, numberOfSeats)
	}

	trainID := adapters.AdaptTrainIDString(TrainIDconst)
	attempt, err := train.BuildReservationAttempt(trainID, 3)
	if err != nil {
		t.Errorf("This should not happen")
	}

	expectedSeatsInAttempt := 3
	seatsInAttempt := len(attempt.Seats())
	if expectedSeatsInAttempt != seatsInAttempt {
		t.Errorf("Expected to have '%d' seats in attempt; got: '%d'",
			expectedSeatsInAttempt, seatsInAttempt)
	}

}
func GetTrainTopology_With_2_Coaches_and_9_seats_are_already_reserved_in_the_first_coach() string {
	return `{"seats": {
		"1A": {"booking_reference": "75bcd15", "seat_number": "1", "coach": "A" },
		"2A": {"booking_reference": "75bcd15", "seat_number": "2", "coach": "A" },
		"3A": {"booking_reference": "75bcd15", "seat_number": "3", "coach": "A" },
		"4A": {"booking_reference": "75bcd15", "seat_number": "4", "coach": "A" },
		"5A": {"booking_reference": "75bcd15", "seat_number": "5", "coach": "A" },
		"6A": {"booking_reference": "75bcd15", "seat_number": "6", "coach": "A" },
		"7A": {"booking_reference": "75bcd15", "seat_number": "7", "coach": "A" },
		"8A": {"booking_reference": "75bcd15", "seat_number": "8", "coach": "A" },
		"9A": {"booking_reference": "75bcd15", "seat_number": "9", "coach": "A" },
		"10A": {"booking_reference": " ", "seat_number": "10", "coach": "A" },
		"1B": {"booking_reference": "", "seat_number": "1", "coach": "B" },
		"2B": {"booking_reference": "", "seat_number": "2", "coach": "B" },
		"3B": {"booking_reference": "", "seat_number": "3", "coach": "B" },
		"4B": {"booking_reference": "", "seat_number": "4", "coach": "B" },
		"5B": {"booking_reference": "", "seat_number": "5", "coach": "B" },
		"6B": {"booking_reference": "", "seat_number": "6", "coach": "B" },
		"7B": {"booking_reference": "", "seat_number": "7", "coach": "B" },
		"8B": {"booking_reference": "", "seat_number": "8", "coach": "B" },
		"9B": {"booking_reference": "", "seat_number": "9", "coach": "B" },
		"10B": {"booking_reference": "", "seat_number": "10", "coach": "B" }
	}}`
}

func Test_Train_should_expose_coaches(t *testing.T) {

	apdatedTrainTopology, err := adapters.AdaptTrainTopology(GetTrainTopology_With_2_Coaches_and_9_seats_are_already_reserved_in_the_first_coach())

	if err != nil {
		t.Fatalf("This should not happen: %s", err.Error())
	}

	train := train.New(apdatedTrainTopology)
	if len(train.Coaches) != 2 {
		t.Errorf("Expected the number of coaches to 2; got %d", len(train.Coaches))
	}

	if len(train.Coaches["A"].Seats()) != 10 {
		t.Errorf("Expected the number of seats in coach A to be 10; got: %d", len(train.Coaches["A"].Seats()))
	}
	if len(train.Coaches["B"].Seats()) != 10 {
		t.Errorf("Expected the number of seats in coach B to be 10; got: %d", len(train.Coaches["B"].Seats()))
	}
	if len(train.Coaches) != 2 {
		t.Errorf("Expected to have 2 coaches; got: %d", len(train.Coaches))
	}
}

func Test_Train_with_6_seats_5_reserverd_should_not_Exceed_the_overall_capacity(t *testing.T) {

	seats := []*train.Seat{
		train.NewSeat("A", 1, train.EmptyBookingReference()),
		train.NewSeat("A", 2, train.EmptyBookingReference()),
		train.NewSeat("A", 3, train.EmptyBookingReference()),
		train.NewSeat("A", 4, train.EmptyBookingReference()),
		train.NewSeat("A", 5, train.EmptyBookingReference()),
		train.NewSeat("A", 6, train.EmptyBookingReference()),
	}
	currentTrain := train.New(seats)
	expectedFalse := currentTrain.DoesNotExceedOveralTrainCapacity(5)
	if expectedFalse == true {
		t.Errorf("expected the result to be 'false'; got: '%t'", expectedFalse)
	}
}

func Test_Failed_ReservationAttempt(t *testing.T) {
	apdatedTrainTopology, err := adapters.AdaptTrainTopology(GetTrainTopology_With_2_Coaches_and_9_seats_are_already_reserved_in_the_first_coach())
	if err != nil {
		t.Errorf("This should not happen!")
	}
	trainID := adapters.AdaptTrainIDString(TrainIDconst)
	currentTrain := train.New(apdatedTrainTopology)
	reservationAttempt, err := currentTrain.BuildReservationAttempt(trainID, 11)
	if reservationAttempt != nil {
		t.Errorf("This should not happen: %+v", reservationAttempt)
	}

}

func Test_Reserved_Seats(t *testing.T) {
	bookingRef := adapters.AdaptBookingReferenceDto(BookingReference)
	seats := []*train.Seat{
		train.NewSeat("A", 1, bookingRef),
		train.NewSeat("A", 2, bookingRef),
		train.NewSeat("A", 3, bookingRef),
		train.NewSeat("A", 4, bookingRef),
		train.NewSeat("A", 5, bookingRef),
		train.NewSeat("A", 6, bookingRef),
		train.NewSeat("A", 7, bookingRef),
		train.NewSeat("A", 8, train.EmptyBookingReference()),
		train.NewSeat("A", 9, train.EmptyBookingReference()),
		train.NewSeat("A", 10, train.EmptyBookingReference()),
	}

	currentTrain := train.New(seats)
	reservedSeats := currentTrain.ReservedSeats()
	if len(reservedSeats) != 7 {
		t.Errorf("Expected the number of reserved seats to be 7; got: %d", reservedSeats)
	}
}
