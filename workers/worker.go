package workers

import (
	"sort"
)

/*Worker runs on a thread and executes a job*/
type Worker struct {
	id      int
	jobs    <-chan Job
	results chan<- *[]int
}

/*NewWorker will initialize a new worker which will run in a separate thread*/
func NewWorker(id int, jobs <-chan Job, results chan<- *[]int) *Worker {
	return &Worker{
		id:      id,
		jobs:    jobs,
		results: results,
	}
}

/*Do will listen on the job channel and send an HTTP request when it finds a job. Result will
be pushed into a separate channel*/
func (worker *Worker) Do() {
	job := <-worker.jobs

	result := GetNumbers(job.Url)

	arr := result.Numbers

	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}

	worker.results <- &arr
}
