package local_io

import (
	"encoding/json"
	"fmt"
	"lygiagretumas_lab2/data_objects"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Parse_data(path string) []data_objects.Book {
	file, err := os.ReadFile(path)
	check(err)

	var books []data_objects.Book
	check(json.Unmarshal(file, &books))

	return books
}
