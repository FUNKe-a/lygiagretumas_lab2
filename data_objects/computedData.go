package data_objects

type ComputedData struct {
	Data *Book
	Hash string
}

var PoisonPillComp = ComputedData{&PoisonPill, "poison"}
