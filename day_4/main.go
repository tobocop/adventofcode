package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	roomNames := strings.Split(string(input), "\n")

	reChecksum := regexp.MustCompile("([a-z][a-z][a-z][a-z][a-z])]$")
	reSectorID := regexp.MustCompile("-([0-9]+)")

	for i := 0; i < len(roomNames); i++ {
		roomName := roomNames[i]
		checksum := reChecksum.FindStringSubmatch(roomName)[1]
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
			var fullCycles float64
			fullCycles = float64(sectorID) / float64(26)
			remainder := sectorID - (int(math.Trunc(fullCycles)) * 26)

			decodedName := ""

			for n := 0; n < len(letters); n++ {
				letter := string(letters[n])
				if letter == "-" {
					decodedName = decodedName + " "
					continue
				}

				newRune := []rune(letter)[0] + rune(remainder)
				// z = 122
				if newRune > 122 {
					newRune = newRune - 26
				}
				decodedName = decodedName + string(newRune)
			}

			if strings.Contains(decodedName, "north") {
				fmt.Println(decodedName)
				fmt.Println("found in sector: ", sectorID)
				break
			}
		}
	}
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
