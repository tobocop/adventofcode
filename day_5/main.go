package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"strconv"
)

func main() {
	var password [8]string
	input := "ugkcyxxp"

	i := 0
	h := md5.New()
	found := 0

	for found < 8 {
		h.Reset()

		thing := input + strconv.Itoa(i)
		hashed := Hash(thing, h)
		if hashed[0:5] == "00000" {
			pos, err := strconv.Atoi(string(hashed[5:6]))
			if err == nil && pos < 8 {
				if password[pos] == "" {
					fmt.Println("match")
					password[pos] = hashed[6:7]
					found++

				}
			}
		}

		i++
	}

	for i := 0; i < 8; i++ {
		fmt.Fprint(os.Stdout, password[i])
	}
	fmt.Println()
}

func Hash(text string, hasher hash.Hash) string {
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
