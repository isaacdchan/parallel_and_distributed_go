package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToIntArray fuck
func StringToIntArray(s string) (ia []int) {
	sa := strings.Split(s, " ")

	ia = make([]int, len(sa))

	for i, val := range sa {
		num, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
			return
		}

		ia[i] = num
	}

	return ia
}

// IntArrayToByteArray fuck
func IntArrayToByteArray(ia []int) (ba []byte) {
	s := ""
	for _, val := range ia {
		s += strconv.Itoa(val) + " "
	}

	s = s[:len(s) - 1]
	s += "\n"
	ba = []byte(s)

	return ba
}

// ByteArrayToIntArray fuck
func ByteArrayToIntArray(ba []byte) (ia []int) {
	ba = ba[:len(ba) -1]
	s := string(ba)
	sa := strings.Split(s, " ")

	ia = make([]int, len(sa))
	for i, val := range sa {
		intVal, err := strconv.Atoi(val)
		if (err == nil) {
			ia[i] = intVal
		}
	}
	
	return ia
}