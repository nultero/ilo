package fn

import (
	"strconv"
	"testing"
)

var datesToCheck = []Day{
	{29, "Jan"},
	{30, "Oct"},
	{1, "Apr"},
}

var mockEventDates = []Day{
	{2, "Feb"},
	{2, "Nov"},
	{4, "Apr"},
}

func TestRollForwardApi(t *testing.T) {

	matches := populateMatchesIndex(len(datesToCheck))

	for i, day := range datesToCheck {
		for z := 0; z < 5; z++ {
			day = RollForwardDay(day)

			if day == mockEventDates[i] {
				matches[i]++
				break
			}
		}
	}

	s := ""
	for i := 0; i < len(datesToCheck); i++ {
		if matches[i] != i+1 {
			s += strconv.Itoa(i) + ","
		}
	}
	s = s[:len(s)-1]
	t.Errorf("TestRollForward failed at case(s): %s", s)
}

// Makes a bare slice of ints for the test cases to increment if successful.
func populateMatchesIndex(numCases int) []int {
	matches := []int{}
	for x := 0; x < len(datesToCheck); x++ {
		matches = append(matches, x)
	}
	return matches
}
