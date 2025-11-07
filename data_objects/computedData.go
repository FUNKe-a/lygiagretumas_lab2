package data_objects

type ComputedData struct {
	Data *Book
	Hash string
}

func PoisonPillComp() ComputedData {
	pill := PoisonPill()
	return ComputedData{&pill, "poison"}
}
