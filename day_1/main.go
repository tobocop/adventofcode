package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var blocks [4]int
	orientation := 0
	input := "R4, R3, R5, L3, L5, R2, L2, R5, L2, R5, R5, R5, R1, R3, L2, L2, L1, R5, L3, R1, L2, R1, L3, L5, L1, R3, L4, R2, R4, L3, L1, R4, L4, R3, L5, L3, R188, R4, L1, R48, L5, R4, R71, R3, L2, R188, L3, R2, L3, R3, L5, L1, R1, L2, L4, L2, R5, L3, R3, R3, R4, L3, L4, R5, L4, L4, R3, R4, L4, R1, L3, L1, L1, R4, R1, L4, R1, L1, L3, R2, L2, R2, L1, R5, R3, R4, L5, R2, R5, L5, R1, R2, L1, L3, R3, R1, R3, L4, R4, L4, L1, R1, L2, L2, L4, R1, L3, R4, L2, R3, L1, L5, R4, R5, R2, R5, R1, R5, R1, R3, L3, L2, L2, L5, R2, L2, R5, R5, L2, R3, L5, R5, L2, R4, R2, L1, R3, L5, R3, R2, R5, L1, R3, L2, R2, R1"

	directions := strings.Split(input, ", ")
	// index where 0=n,1=e,2=s,3=w
	type coord [2]int
	standingAt := coord{0, 0}
	xy := map[coord]int{standingAt: 1}

	done := false

	for i := 0; i < len(directions); i++ {
		d := directions[i]
		way := string([]rune(d)[0])
		a := strings.TrimPrefix(d, "R")
		a = strings.TrimPrefix(a, "L")
		amount, _ := strconv.Atoi(a)

		if way == "R" {
			orientation++
		} else {
			orientation = orientation - 1
		}

		if orientation > 3 {
			orientation = 0
		}

		if orientation < 0 {
			orientation = 3
		}

		for n := 0; n < amount; n++ {

			blocks[orientation] = blocks[orientation] + 1

			switch orientation {
			case 0:
				standingAt[1] = standingAt[1] + 1
			case 1:
				standingAt[0] = standingAt[0] + 1
			case 2:
				standingAt[1] = standingAt[1] - 1
			case 3:
				standingAt[0] = standingAt[0] - 1
			}

			xy[standingAt] = xy[standingAt] + 1
			if xy[standingAt] == 2 {
				done = true
				break
			}
		}

		if done {
			break
		}
	}

	ns := math.Abs(float64((blocks[0] - blocks[2])))
	ew := math.Abs(float64((blocks[1] - blocks[3])))

	fmt.Println(ns + ew)
}
