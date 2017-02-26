package numbers

import (
	"log"
	"net/url"
	"number-service/workers"
	"time"
)

/*Get takes an array of URL's and a timer. It returns the result of
merging the numbers returned by each of those URL's in sorted order*/
func Get(urls []string, timer *time.Timer) *[]int {
	jobs := getJobs(urls)

	results := workers.NewWorkerPool(jobs, len(jobs)).Process()

	mergedArray := NewAccumulator(results, len(jobs)).Accumulate(timer)

	return mergedArray
}

func getJobs(urls []string) []workers.Job {
	jobs := []workers.Job{}

	for _, u := range urls {
		validatedUrl, err := url.ParseRequestURI(u)

		if err == nil {
			j := workers.Job{
				Url: validatedUrl.String(),
			}

			jobs = append(jobs, j)
		} else {
			log.Printf("URL validation error : %s", err.Error())
		}
	}

	return jobs
}
