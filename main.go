package main

import (
	// "encoding/json"
	"fmt"
	// "lygiagretumas_lab2/data_objects"
	"lygiagretumas_lab2/local_io"
)

func main() {
	books := local_io.Parse_data("data/IFF-310_KucinskasR_L1b_dat_1.json")

	for _, u := range books {
		fmt.Println(u)
	}
}
