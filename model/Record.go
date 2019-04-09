package model

type ReocrdList struct {
	data []Reocrd
}
type Reocrd struct {
	spo2 int32
	dpm  int32
	time int64
}
