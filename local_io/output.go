package local_io

import (
	"fmt"
	"lygiagretumas_lab2/data_objects"
	"os"
	"strings"
)

func OutputToFile(path string, data []*data_objects.ComputedData) {
	f, err := os.Create(path)
	check(err)

	defer f.Close()

	fmt.Fprintf(f, "%s\n", strings.Repeat("-", 98))
	fmt.Fprintf(f, "|%-16s|%-7s|%-5s|%-65s|\n", "isbn", "price", "count", "hash")
	fmt.Fprintf(f, "%s\n", strings.Repeat("-", 98))

	for _, v := range data {
		fmt.Fprintf(f, "|%-16s|%-7.2f|%-5d|%-65s|\n", v.Data.Isbn, v.Data.Price, v.Data.Count, v.Hash)
	}
	fmt.Fprintf(f, "%s\n", strings.Repeat("-", 98))
	f.Sync()
}
