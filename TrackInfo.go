package main

import (
	"time"
)

type TrackInfo struct {
	MbId             string
	Name             string
	EarliestPlayDate time.Time
	PlayCount        int
}

func (t TrackInfo) GetNumberOfDaysSinceEarliestPlayDate() int {
	return int(time.Now().Sub(t.EarliestPlayDate).Hours() / 24)
}

func (t TrackInfo) GetPlayFrequency() float32 {
	return float32(t.PlayCount) / float32(t.GetNumberOfDaysSinceEarliestPlayDate())
}

type TrackInfoList []TrackInfo

func (s TrackInfoList) Len() int {
	return len(s)
}

func (s TrackInfoList) Less(i, j int) bool {
	return s[i].GetPlayFrequency() < s[j].GetPlayFrequency()
}

func (s TrackInfoList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
