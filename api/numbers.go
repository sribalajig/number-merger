package api

import (
	"encoding/json"
	"net/http"
	"number-service/numbers"
	"time"
)

/*Get returns the numbers as per the request*/
func Get() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := time.NewTimer(time.Millisecond * 480) // start a timer

		queryParams := r.URL.Query()

		if urls, ok := queryParams["u"]; ok && len(urls) > 0 {
			nums := numbers.Get(urls, timer) // call the service with url's and the timer which is set to 480ms

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(map[string]interface{}{"Numbers": nums})
		} else { // there are no query params for numbers, so we say that its a bad request
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
