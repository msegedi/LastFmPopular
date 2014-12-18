package main

import "encoding/json"
import "io/ioutil"
import "fmt"
import "os"

type PlayTrack struct {
	MbId string
}

type Play struct {
	PlayTrack PlayTrack
}

type TrackInfo struct {
	MbId      string
	PlayCount int
}

func main() {
	jsonData, e := ioutil.ReadFile("./test.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	//fmt.Printf("%s\n", string(file))
	//fmt.Println("hello world!")

	var plays []Play
	err := json.Unmarshal(jsonData, &plays)
	if err != nil {
		fmt.Println("error reading JSON data")
		os.Exit(1)
	}

	tracks := make(map[string]int)
	for _, play := range plays {
		tracks[play.PlayTrack.MbId]++
	}
	fmt.Println(tracks)
}
