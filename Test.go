package main

import "encoding/json"
import "io/ioutil"
import "fmt"
import "os"
import "time"
import "sort"

// *** Last.Fm Classes ***
type Play struct {
	Timestamp Timestamp
	Track     Track
}

type Timestamp struct {
	//Iso time.Time
	UnixTimestamp int64
}

type Track struct {
	MbId string
	Name string
}

// *** App Classes ***
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

func main() {
	// Build a list of plays for all JSON data in the History folder.
	plays := getAllPlaysForFolder("./History")

	// Create a track info list with details about each track in the play history.
	tracks := make(TrackInfoList, 0)
	for _, play := range plays {
		i := getIndexOfTrack(tracks, play.Track.MbId)
		if i >= 0 {
			tracks[i].PlayCount++
			if time.Unix(play.Timestamp.UnixTimestamp, 0).Before(tracks[i].EarliestPlayDate) {
				tracks[i].EarliestPlayDate = time.Unix(play.Timestamp.UnixTimestamp, 0)
			}
		} else {
			tracks = append(tracks, TrackInfo{MbId: play.Track.MbId, Name: play.Track.Name, EarliestPlayDate: time.Unix(play.Timestamp.UnixTimestamp, 0), PlayCount: 1})
		}
	}

	// Sort the track info list by track play frequency.
	sort.Sort(tracks)

	// Display the results.
	for _, track := range tracks {
		fmt.Printf("Frequency: %12g | Plays: %6d | Song: %s\n", track.GetPlayFrequency(), track.PlayCount, track.Name)
	}
}

func getAllPlaysForFolder(folderName string) []Play {
	var plays []Play

	files, _ := ioutil.ReadDir(folderName)
	for _, f := range files {
		// Read in JSON data.
		jsonData, e := ioutil.ReadFile(folderName + "/" + f.Name())
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		}

		// Move JSON data into list of plays.
		var filePlays []Play
		err := json.Unmarshal(jsonData, &filePlays)
		if err != nil {
			fmt.Printf("error reading JSON data insert for file: %s\nError message: %s\n", f.Name(), err)
			os.Exit(1)
		}

		plays = append(plays, filePlays...)
	}

	return plays
}

// ToDo: Have this function return a track instead of an index.
func getIndexOfTrack(tracks []TrackInfo, mbId string) int {
	for i, track := range tracks {
		if track.MbId == mbId {
			return i
		}
	}
	return -1
}
