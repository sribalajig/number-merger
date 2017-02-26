package workers

import (
	"encoding/json"
	"log"
	"net/http"
)

/*GetNumbers executes a HTTP request. It will return an empty
array if there is an error or the status code is not 200*/
func GetNumbers(url string) *Result {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error while retreiving numbers from remote API %s : %s", url, err.Error())

		return &Result{
			Numbers: []int{},
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Remote API %s returned error %d", url, resp.StatusCode)

		return &Result{
			Numbers: []int{},
		}
	}

	result := new(Result)

	json.NewDecoder(resp.Body).Decode(result)

	return result
}
