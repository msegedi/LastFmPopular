package main

type Play struct {
	Timestamp Timestamp
	Track     Track
}

type Timestamp struct {
	UnixTimestamp int64
}

type Track struct {
	MbId string
	Name string
}
