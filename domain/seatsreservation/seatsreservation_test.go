package seatsreservation_test

import (
	"testing"
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

func Test_Reserve_seats_when_train_is_emptry(t *testing.T) {

	trainID := "hello_id"
	bookingReference := "75bcd15"
	numberOfSeatsToRequest := 3

	ticketManager := webticketmanager.New()

	jsonReservation := ticketManager.Reserve(trainID, numberOfSeatsToRequest)

	epectedJson := `{"train_id": "hello_id", "booking_reference": "75bcd15","seats": ["1A","2A","3A"]}`

	if jsonReservation != epectedJson {
		t.Errorf("Expected json reservation to be: %s; but got: %s", epectedJson, jsonReservation)
	}
}

func Test_Not_Reserve_Seats_when_it_exceeds_max_capacity_threshold(t *testing.T) {

	numberOfSeatsToRequest := 3

	expectedJson := `{"train_id": "hello_id", "booking_reference": "", "seats": []}`

	jsonReservation := ``
	if jsonReservation != expectedJson {
		t.Errorf("Expected json reservation to be: %s; but got: %s", expectedJson, jsonReservation)
	}

}

func Test_Reserve_all_seats_in_the_same_coach(t *testing.T) {
	numberOfSeatsToRequest := 2

	jsonReservation := ``
	expectedJson := `{"train_id":"hello_id", "booking_reference": "75bcd15", "seats": ["1B","2B"] }`

	if jsonReservation != expectedJson {
		t.Errorf("Expected json reservation to be: %s; but got: %s", expectedJson, jsonReservation)
	}
}
