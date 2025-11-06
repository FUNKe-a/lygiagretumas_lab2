package main

import (
	// "encoding/json"
	"lygiagretumas_lab2/data_objects"
	"lygiagretumas_lab2/local_io"
)

func main() {
	books := local_io.Parse_data("data/IFF-310_KucinskasR_L1b_dat_1.json")
	main_data_ch := make(chan data_objects.Book)
	worker_data_ch := make(chan data_objects.Book)

	go dataThread(len(books), main_data_ch, worker_data_ch)

	for _, v := range books {
		main_data_ch <- v
	}
	close(main_data_ch)
}

func dataThread(size int, add <-chan data_objects.Book, remove chan data_objects.Book) {
	books := make([]data_objects.Book, size)
	i := 0

	for {
		var addChan <-chan data_objects.Book
		var remChan chan data_objects.Book
		if i < size {
			addChan = add
		}
		if i > 0 {
			remChan = remove
		}

		select {
		case book := <-addChan:
			books[i] = book
			i++
		case request := <-remChan:
			if request == data_objects.Request() {
				remove <- books[i]
				i--
			}
		}
	}
}

func workerThread(size int, request chan data_objects.Book, result <-chan data_objects.Book) {

}
