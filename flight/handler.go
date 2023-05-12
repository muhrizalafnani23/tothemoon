package flight

import (
	"encoding/json"
	"net/http"
)

func FlightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Query().Get("code") != "" {
			res, err := GetFlightByCode(r.URL.Query().Get("code"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			resByte, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(resByte)
			return
		}

		res := GetListFlight()
		resByte, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resByte)
		return
	}

	if r.Method == "POST" {
		// baca body
		var flg Flight
		err := json.NewDecoder(r.Body).Decode(&flg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if flg.FlightCode == "" {
			http.Error(w, "data not complete", http.StatusBadRequest)
			return
		}

		// panggil logic
		res, err := InsertNewFlight(flg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// response client
		resJson, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(resJson)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
