package workers

/*WorkerPool will start a take an array of jobs and run each of them in a worker thread*/
type WorkerPool struct {
	jobs []Job
	size int
}

/*NewWorkerPool takes an array of jobs and the size of the pool*/
func NewWorkerPool(jobs []Job, size int) *WorkerPool {
	return &WorkerPool{
		jobs: jobs,
		size: size,
	}
}

/*Process will start the worker pool - the number of threads equals the size of the pool*/
func (workerPool WorkerPool) Process() chan *[]int {
	results := make(chan *[]int, len(workerPool.jobs))
	workerJobs := make(chan Job, len(workerPool.jobs))

	for i := 0; i <= workerPool.size; i++ {
		worker := NewWorker(i, workerJobs, results)

		go worker.Do()
	}

	for _, job := range workerPool.jobs {
		workerJobs <- job
	}

	return results
}
