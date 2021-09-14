package fn

import (
	"bufio"
	"os"
	"strings"
)

func IsEmpty(s interface{}) bool {

	t, ok := s.(string)
	if ok {
		if len(t) == 0 {
			return true
		}
	}

	sl, ok := s.([]string)
	if ok {
		if len(sl) == 0 {
			return true
		}
	}

	return false
}

func GetInput() string {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	return strings.ReplaceAll(s, "\n", "")
}
