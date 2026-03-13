package main

import (
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	fmt.Println(Encode(62))
}

func Encode(id uint64) string {
	length := uint64(len(alphabet))
	var sbuilder strings.Builder
	for id > 0 {
		remainder := id % length
		sbuilder.WriteByte(alphabet[remainder])
		id /= length
	}
	res := sbuilder.String()
	runeSlice := []rune(res)
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	return string(runeSlice)
}
