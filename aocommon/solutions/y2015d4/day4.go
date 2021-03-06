package y2015d4

import (
	"advent/aocommon/solutions"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year:         15,
	Day:          4,
	DefaultInput: "bgvyzdsv",
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(in string) string {
	return strconv.Itoa(getKeyForInput(in, "00000"))
}

func a2(in string) string {
	return strconv.Itoa(getKeyForInput(in, "000000"))
}

//This is probably the slowest possible way to do this in go, but oh well.
//Santa will just end up getting an ASIC for mining anyways
func getKeyForInput(i string, searchString string) int {
	counter := 0
	for true {
		testString := i + strconv.Itoa(counter)
		out := md5.Sum([]byte(testString))
		testOut := fmt.Sprintf("%x", out)
		if strings.HasPrefix(testOut, searchString) {
			return counter
		}
		counter++
	}
	return 0
}
