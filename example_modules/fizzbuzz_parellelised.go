package examplemodules

import(
	"log"
)


func Worker_pool_fizzbuzz(start, end, worker_count int) [][]string {

	// use a worker pool to orchestrate the parallelization of Fizzbuzz

	jump_div_check := (end - start) % worker_count
	if jump_div_check != 0 {
		log.Fatalln("The degree must perfectly divide the number of entries")
	}

	jump := (end - start) / worker_count
	iter_start := start

	job_channel := make(chan []int, worker_count)
	results_channel := make(chan []string, worker_count)

	var results [][]string

	for i := 0; i < worker_count; i++ {

		go func(job_chan <-chan []int, result_chan chan<- []string) {
			// Worker!
			for start_end_stats := range job_chan {
				// Spawn the new job
				start := start_end_stats[0]
				end := start_end_stats[1]
				// Send the results back to the
				result_chan <- Fizzbuzz_sequential(start, end)
			}
		}(job_channel, results_channel)

	}

	for {
		if iter_start < end {
			range_start := iter_start
			range_end := iter_start + jump
			// fmt.Println("Dispatching Job (START, END)", range_start+1, range_end)
			job_channel <- []int{range_start + 1, range_end}
			iter_start += jump
		} else {
			break
		}
	}

	close(job_channel)

	for i := 0; i < worker_count; i++ {
		results = append(results, <-results_channel)
	}
	return results

}
