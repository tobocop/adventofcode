package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//411182
func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic("NOPE")
	}

	// input := "totally-real-room-200[decoy]"
	roomNames := strings.Split(string(input), "\n")

	re := regexp.MustCompile("([a-z][a-z][a-z][a-z][a-z])]$")
	reSectorID := regexp.MustCompile("-([0-9]+)")
	sumIDs := 0

	for i := 0; i < len(roomNames); i++ {
		roomName := roomNames[i]
		checksum := re.FindStringSubmatch(roomName)[1]
		sectorID, _ := strconv.Atoi(reSectorID.FindStringSubmatch(roomName)[1])

		letters := roomName[0:reSectorID.FindStringIndex(roomName)[0]]
		validationThingy := validationThingy{}

		for n := 0; n < len(checksum); n++ {
			checksumLetter := string(checksum[n])
			checksumLetterCount := strings.Count(letters, checksumLetter)
			validationThingy = append(validationThingy, thingyType{
				Letter: checksumLetter,
				Count:  checksumLetterCount,
			})
		}

		if validationThingy.Valid() {
			sumIDs = sumIDs + sectorID
		}

		// fmt.Println("******************************")
		// fmt.Println(roomName)
		// fmt.Println(checksum)
		// fmt.Println(sectorID)
		// fmt.Println(letters)
		// fmt.Println(validationThingy.Valid())
		// fmt.Println("******************************")
	}
	fmt.Println(sumIDs)
}

type thingyType struct {
	Letter string
	Count  int
}

func (tt thingyType) RuneVal() int {
	return int([]rune(tt.Letter)[0])
}

type validationThingy []thingyType

func (vt validationThingy) Valid() bool {
	for i := 0; i < len(vt)-1; i++ {
		c := vt[i]
		n := vt[i+1]
		if c.Count < n.Count {
			return false
		}

		if c.Count == 0 || n.Count == 0 {
			return false
		}

		if c.Count == n.Count && c.RuneVal() > n.RuneVal() {
			return false
		}
	}

	return true
}
