package main

import (
	"container/list"
	// "crypto/sha256"
	// "encoding/hex"
	// "fmt"
	"log"
	"lygiagretumas_lab2/data_objects"
	"lygiagretumas_lab2/local_io"
)

const threadCount = 1

func main() {
	books := local_io.Parse_data("data/IFF-310_KucinskasR_L1b_dat_1.json")
	main_data_ch := make(chan *data_objects.Book)
	worker_data_ch := make(chan bool)
	data_worker_ch := make(chan *data_objects.Book)
	worker_result_ch := make(chan *data_objects.ComputedData)
	result_main_ch := make(chan []data_objects.Book)

	go dataThread(len(books)/2, main_data_ch, worker_data_ch, data_worker_ch)

	for range threadCount {
		go workerThread(worker_data_ch, data_worker_ch, worker_result_ch)
	}

	for _, v := range books {
		main_data_ch <- &v
	}

	for range threadCount {
		pill := data_objects.PoisonPill()
		main_data_ch <- &pill
	}

	close(main_data_ch)

	results := <-result_main_ch

	log.Println(results)
}

func dataThread(size int, add <-chan *data_objects.Book, request <-chan bool, send chan<- *data_objects.Book) {
	books := list.New()
	poison_count := 0

	for poison_count < threadCount {
		var addChan <-chan *data_objects.Book
		var reqChan <-chan bool
		if books.Len() < size {
			addChan = add
		}
		if books.Len() > 0 {
			reqChan = request
		}

		select {
		case book := <-addChan:
			books.PushBack(book)

		case request := <-reqChan:
			if request == true {
				element := books.Remove(books.Front()).(*data_objects.Book)
				if *element == data_objects.PoisonPill() {
					poison_count++
				}
				send <- element
			}
		}

	}
	close(send)
}

func workerThread(request chan<- bool, receive <-chan *data_objects.Book, send chan<- *data_objects.ComputedData) {
	for {
		request <- true
		value := <-receive

		if *value == data_objects.PoisonPill() {
			//pill := data_objects.PoisonPillComp()
			//send <- &pill
			break
		}

		//s := fmt.Sprintf("%d|%.2f|%d", value.Isbn, value.Price, value.Count)
		//h := sha256.Sum256([]byte(s))
		//result := data_objects.ComputedData{value, hex.EncodeToString(h[:])}
		//send <- &result
	}
	close(request)
}
