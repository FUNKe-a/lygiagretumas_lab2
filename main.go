package main

import (
	"container/list"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"lygiagretumas_lab2/data_objects"
	"lygiagretumas_lab2/local_io"
	// "sort"
	"time"
	"unicode"
	"unicode/utf8"
)

const threadCount = 10

func main() {
	books := local_io.Parse_data("data/IFF-310_KucinskasR_L1b_dat_1.json")
	main_data_ch := make(chan *data_objects.Book)
	worker_data_ch := make(chan bool)
	data_worker_ch := make(chan *data_objects.Book)
	worker_result_ch := make(chan *data_objects.ComputedData)
	// result_main_ch := make(chan []*data_objects.ComputedData)

	go dataThread(len(books)/2, main_data_ch, worker_data_ch, data_worker_ch)

	for range threadCount {
		go workerThread(worker_data_ch, data_worker_ch, worker_result_ch)
	}

	// go resultThread(len(books)+1, worker_result_ch, result_main_ch)

	for _, v := range books {
		main_data_ch <- &v
	}

	for range threadCount {
		main_data_ch <- &data_objects.PoisonPill
	}

	close(main_data_ch)

	time.Sleep(time.Second * 2)

	// results := <-result_main_ch
	//
	// for _, v := range results {
	// 	log.Println(*v)
	// }
}

func dataThread(size int, add <-chan *data_objects.Book, request chan bool, send chan<- *data_objects.Book) {
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

		case <-reqChan:
			element := books.Remove(books.Front()).(*data_objects.Book)
			if element == &data_objects.PoisonPill {
				poison_count++
			}
			send <- element
		}

	}
	close(request)
	close(send)
}

func workerThread(request chan<- bool, receive <-chan *data_objects.Book, send chan<- *data_objects.ComputedData) {
	for {
		request <- true
		value := <-receive

		log.Println(value)
		if value == &data_objects.PoisonPill {
			// send <- &data_objects.PoisonPillComp
			break
		}

		s := fmt.Sprintf("%s|%.2f|%d", value.Isbn, value.Price, value.Count)
		h := sha256.Sum256([]byte(s))
		result := data_objects.ComputedData{Data: value, Hash: hex.EncodeToString(h[:])}
		rune1, _ := utf8.DecodeRuneInString(result.Hash)
		if unicode.IsLetter(rune1) {
			// send <- &result
		}
	}
}

// func resultThread(capacity int, receive <-chan *data_objects.ComputedData, send chan<- []*data_objects.ComputedData) {
// 	data := make([]*data_objects.ComputedData, 0, capacity)
// 	n := 0
//
// 	for {
// 		element := <-receive
// 		if element == &data_objects.PoisonPillComp {
// 			break
// 		}
// 		i := sort.Search(len(data), func(j int) bool {
// 			return data[j].Hash >= element.Hash
// 		})
// 		data = append(data, nil)
// 		copy(data[i+1:], data[i:])
// 		data[i] = element
// 		n++
// 	}
//
// 	send <- data
// 	close(send)
// }
