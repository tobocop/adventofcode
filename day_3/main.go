package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic("NOPE")
	}

	t := strings.Split(string(input), "\n")

	goodOnes := 0
	sides := []int{}

	for i := 0; i < len(t); i++ {

		s := strings.Split(t[i], " ")
		for n := 0; n < len(s); n++ {

			d, err := strconv.Atoi(s[n])
			if err != nil {
				continue
			}
			sides = append(sides, d)
		}

		if len(sides) == 9 {
			newSet := [][]int{
				{sides[0], sides[3], sides[6]},
				{sides[1], sides[4], sides[7]},
				{sides[2], sides[5], sides[8]},
			}

			for d := 0; d < len(newSet); d++ {
				sort.Ints(newSet[d])
				if (newSet[d][0] + newSet[d][1]) > newSet[d][2] {
					goodOnes++
				}
			}

			sides = []int{}
		}

	}

	fmt.Println(goodOnes)
}
