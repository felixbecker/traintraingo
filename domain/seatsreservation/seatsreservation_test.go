package seatsreservation_test

import (
	"fmt"
	"testing"
	"traintraingo/domain/seatsreservation"
	"traintraingo/infra"
)

func GetTrainTopologyWith10AvailableSeats() string {
	return `{"seats": {
		"1A": {"booking_reference": "", "seat_number": "1", "coach": "A" },
		"2A": {"booking_reference": "", "seat_number": "2", "coach": "A" },
		"3A": {"booking_reference": "", "seat_number": "3", "coach": "A" },
		"4A": {"booking_reference": "", "seat_number": "4", "coach": "A" },
		"5A": {"booking_reference": "", "seat_number": "5", "coach": "A" },
		"6A": {"booking_reference": "", "seat_number": "6", "coach": "A" },
		"7A": {"booking_reference": "", "seat_number": "7", "coach": "A" },
		"8A": {"booking_reference": "", "seat_number": "8", "coach": "A" },
		"9A": {"booking_reference": "", "seat_number": "9", "coach": "A" },
		"10A": {"booking_reference": "", "seat_number": "10", "coach": "A" }
	}`
}

func Test_Reserve_seats_when_train_is_empty(t *testing.T) {
	const requestedSeatCount int = 3

	// Step1: Instantiate the "I want to go out" adapters
	trainDataServiceAdapter := &TrainDataServiceMock{
		JsonResponseString: GetTrainTopologyWith10AvailableSeats(),
	}

	bookingReferenceServiceAdapter := &BookingReferenceServiceMock{
		BookingReference: BookingReference,
	}

	hexagon := seatsreservation.New(trainDataServiceAdapter, bookingReferenceServiceAdapter)

	reservationRequestDto := infra.ReservationRequestDto{
		TrainID:       TrainIDconst,
		NumberOfSeats: requestedSeatCount,
	}

	seatReservationAdapter := infra.NewSeatReservationAdapter(hexagon)

	jsonBytes := seatReservationAdapter.Post(reservationRequestDto)

	expectedJsonString := fmt.Sprintf(`{"train_id":"%s","booking_reference":"%s","seats":["1A","2A","3A"]}`,
		TrainIDconst, BookingReference)
	jsonString := string(jsonBytes)

	if jsonString != expectedJsonString {
		t.Errorf("Expected the response to be %s; got: %s", expectedJsonString, jsonString)
	}

}
func Test_Not_reserve_seats_when_it_exceed_max_capacty_threshold(t *testing.T) {

	const requestedSeatCount int = 3

	// Step1: Instantiate the "I want to go out" adapters
	trainDataServiceAdapter := &TrainDataServiceMock{
		JsonResponseString: GetTrainTopologyWith_10_seats_and_6_already_reserved(),
	}

	bookingReferenceServiceAdapter := &BookingReferenceServiceMock{
		BookingReference: BookingReference,
	}

	hexagon := seatsreservation.New(trainDataServiceAdapter, bookingReferenceServiceAdapter)

	seatReservationAdapter := infra.NewSeatReservationAdapter(hexagon)

	reservationRequestDto := infra.ReservationRequestDto{
		TrainID:       TrainIDconst,
		NumberOfSeats: requestedSeatCount,
	}
	jsonBytes := seatReservationAdapter.Post(reservationRequestDto)
	expectedJsonString := fmt.Sprintf(`{"train_id":"%s","booking_reference":"","seats":[]}`,
		TrainIDconst)
	jsonString := string(jsonBytes)

	if jsonString != expectedJsonString {
		t.Errorf("Expected the response to be %s; got: %s", expectedJsonString, jsonString)
	}
}

func Test_Reserve_all_seats_in_the_same_coach(t *testing.T) {
	const requestedSeatCount int = 2

	// Step1: Instantiate the "I want to go out" adapters
	trainDataServiceAdapter := &TrainDataServiceMock{
		JsonResponseString: GetTrainTopology_With_2_Coaches_and_9_seats_are_already_reserved_in_the_first_coach(),
	}

	bookingReferenceServiceAdapter := &BookingReferenceServiceMock{
		BookingReference: BookingReference,
	}

	hexagon := seatsreservation.New(trainDataServiceAdapter, bookingReferenceServiceAdapter)

	seatReservationAdapter := infra.NewSeatReservationAdapter(hexagon)

	reservationRequestDto := infra.ReservationRequestDto{
		TrainID:       TrainIDconst,
		NumberOfSeats: requestedSeatCount,
	}
	jsonBytes := seatReservationAdapter.Post(reservationRequestDto)
	expectedJsonString := fmt.Sprintf(`{"train_id":"%s","booking_reference":"%s","seats":["1B","2B"]}`,
		TrainIDconst, BookingReference)
	jsonString := string(jsonBytes)

	if jsonString != expectedJsonString {
		t.Errorf("Expected the response to be %s; got: %s", expectedJsonString, jsonString)
	}
}
