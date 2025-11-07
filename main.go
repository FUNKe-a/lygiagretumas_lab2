package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"lygiagretumas_lab2/data_objects"
	"lygiagretumas_lab2/local_io"
)

func main() {
	books := local_io.Parse_data("data/IFF-310_KucinskasR_L1b_dat_1.json")
	main_data_ch := make(chan data_objects.Book)
	// worker_result_ch := make(chan data_objects.Book)
	worker_data_ch := make(chan bool)
	data_worker_ch := make(chan data_objects.Book)

	go dataThread(len(books)/2, len(books), main_data_ch, worker_data_ch, data_worker_ch)
	go workerThread(worker_data_ch, data_worker_ch)

	for _, v := range books {
		main_data_ch <- v
	}
	close(main_data_ch)
}

func dataThread(size int, el_count int, add <-chan data_objects.Book, request <-chan bool, send chan<- data_objects.Book) {
	books := make([]data_objects.Book, size)
	i := 0

	for range el_count * 2 {
		var addChan <-chan data_objects.Book
		var reqChan <-chan bool
		if i < size {
			addChan = add
		}
		if i > 0 {
			reqChan = request
		}
		select {
		case book := <-addChan:
			books[i] = book
			i++
		case request := <-reqChan:
			if request == true {
				i--
				send <- books[i]
			}
		}
	}
	close(send)
}

func workerThread(request chan<- bool, receive <-chan data_objects.Book) {
	for {
		request <- true
		value, ok := <-receive
		if !ok {
			break
		}
		s := fmt.Sprintf("%d|%.2f|%d", value.Isbn, value.Price, value.Count)
		h := sha256.Sum256([]byte(s))
		result := data_objects.ComputedData{value, hex.EncodeToString(h[:])}
		log.Println(result)
	}
	close(request)
}
