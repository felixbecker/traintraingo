package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"traintraingo/infra"

	"github.com/gorilla/mux"
)

func handleReservation(reservationService infra.SeatReservationAdapter) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var dto infra.ReservationRequestDto
		err := decoder.Decode(&dto)

		if err != nil {
			http.Error(w, "Can not parse reservation request", http.StatusBadRequest)
		}

		bytes, err := reservationService.Post(dto)
		if err != nil {
			http.Error(w, fmt.Sprintf("Reservation failed: %s", err.Error()), http.StatusBadRequest)
		}
		w.Write(bytes)
	}
}

func Routes(reservationService infra.SeatReservationAdapter) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/reservations", handleReservation(reservationService))
	return r
}
