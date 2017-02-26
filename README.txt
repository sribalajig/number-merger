How to run the application -

- You need to start the HTTP server by running main.go. Example -> go run main.go -http.addr ":8081"
   By default, the server will run on port 8080.

- A valid request will be of the form -> http://localhost:8080/numbers?u=http://localhost:8090/rand&u=http://localhost:8090/primes

- You can run the unit tests by navigating to the root and executing -> go test ./...

Rationale behind the program -

- Once a request is received, a timer is started.

- After the timer is started, a worker pool is initiated. A worker pool consists of workers, each of which is a Go routine.

- A worker is assigned a job - in this case, a job consists of a URL from which numbers have to be retrieved. A worker runs in a thread of its own (it is  a Go routine). The number of workers in the pool is currently set to the number of jobs. We can change this later on once we determine how many workers are required heuristically.

- Once a worker completes its assigned job, it checks if the resulting array of integers is sorted. If it is not sorted, it will sort the array and write the result on to a buffered channel.

- If the worker encounters an error, it will write an empty array into the buffered channel.

- There is an accumulator which listens on the buffered channel to check for results. The job of the accumulator is to consolidate results which have been returned by the workers. When a result arrives, it merges the result to an array which it maintains. The merging happens in a separate go routine.

- As the accumulator iterates over the buffered channel, it checks if there has been a timeout. If a timeout occurs, we return the result that we have computed till now.

- The accumulator uses a Merge function which runs in linear time.

Ambiguities -

- If a time out occurs while the merge is taking place, I return whatever the latest merged array is. For example, if we have three arrays A, B, C and we merged A, B but a timeout occured while merging C, we will return the result of merging A, B.