package domain_test

import (
	"testing"
	"traintraingo/domain"
	"traintraingo/infra/adapters"
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

	train := domain.NewTrain(apdatedTrainTopology)
	if len(train.Coaches) != 2 {
		t.Errorf("Expected the number of coaches to 2; got %d", len(train.Coaches))
	}

	if len(train.Coaches["A"].Seats()) != 10 {
		t.Errorf("Expected the number of seats in coach A to be 10; got: %d", len(train.Coaches["A"].Seats()))
	}
	if len(train.Coaches["B"].Seats()) != 10 {
		t.Errorf("Expected the number of seats in coach B to be 10; got: %d", len(train.Coaches["B"].Seats()))
	}
}
