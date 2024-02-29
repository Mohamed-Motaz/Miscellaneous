package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("hi")
}

// This is the only func that deals with multi-byte characters
func ConcatStringsPlusEqual(toConcat string) string {
	res := ""
	for i := 0; i < 1000; i++ {
		res += toConcat
	}
	return res
}

func ConcatStringsPlusEqualImproved(toConcat string) string {
	sz := len(toConcat)
	res := make([]byte, 1000*sz, 1000*sz)

	for i := 0; i < 1000; i++ {
		for j := 0; j < sz; j++ {
			res[i*sz+j] = toConcat[j]
		}

	}
	return string(res)
}

func ConcatStringsBytesBuffer(toConcat string) string {
	res := bytes.Buffer{}
	for i := 0; i < 1000; i++ {
		res.WriteString(toConcat)
	}
	return res.String()
}

func ConcatStringsBytesBufferImproved(toConcat string) string {
	res := bytes.Buffer{}
	res.Grow(1000 * len(toConcat))

	for i := 0; i < 1000; i++ {
		res.WriteString(toConcat)
	}
	return res.String()
}
