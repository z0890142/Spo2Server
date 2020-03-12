package model

type ReocrdList struct {
	data []Reocrd
}
type Reocrd struct {
	spo2 int32
	dpm  int32
	time int64
}

type InsertTag struct {
	Start int64
	End   int64
	Tag   int
}
