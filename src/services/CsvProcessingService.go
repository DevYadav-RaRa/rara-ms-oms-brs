package services

import (
	"log"
	"sync"

	"github.com/RaRa-Delivery/rara-ms-boilerplate/src/models"
)

var wg sync.WaitGroup

func worker(input chan int64, output chan int64, uploadId string, limit int64) {
	defer wg.Done()
	// Consumer: Process items from the input channel and send results to output channel
	log.Println("Consumer: Process items from the input channel and send results to output channel")
	for value := range input {
		orderList, _, error := models.FetchOrdersFromDB(uploadId, value, limit)
		if !error {
			log.Println("TODO: process these orders through queues for further processing")
			//TODO: process these orders through queues for further processing
			output <- int64(len(orderList))
		}

	}
}

func Processing(uploadId string, limit int64) {
	log.Println("I am inside Processing with UploadId: ", uploadId)
	_, paginationData, err := models.FetchOrdersFromDB(uploadId, 1, limit)
	log.Println("FetchOrdersFromDB err: ", err)
	if !err {
		log.Println("paginationData: ", paginationData)
		input := make(chan int64, paginationData.TotalPage)
		output := make(chan int64, paginationData.TotalPage)
		workers := paginationData.TotalPage
		//workers := int64(250)

		for i := int64(1); i <= workers; i++ {
			wg.Add(1)
			go worker(input, output, uploadId, limit)

			log.Println("Producer: load up input channel with pagination index objects")
			// Producer: load up input channel with pagination index objects
			input <- i

		}

		// Close input channel since no more orderList are being sent to input channel
		log.Println("Close input channel since no more orderList are being sent to input channel")
		close(input)
		// Wait for all goroutines to finish processing
		log.Println("Wait for all goroutines to finish processing")
		wg.Wait()
		log.Println("Close output channel since all workers have finished processing")
		// Close output channel since all workers have finished processing
		close(output)
		log.Println("Read from output channel")
		// Read from output channel
		for result := range output {
			log.Println("Result: ", result)
		}

	}
}
