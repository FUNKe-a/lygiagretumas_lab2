package data_objects

type Book struct {
	Isbn  string  `json:"isbn"`
	Price float32 `json:"price"`
	Count uint    `json:"count"`
}
