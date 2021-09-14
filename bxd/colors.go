package bxd

import "fmt"

// Returns blued ASCII
func Blue(s string) string {
	return fmt.Sprintf("\x1b[32;1;4m%v\x1b[0m", s)
}

// Returns bolded ASCII
func Emph(s string) string {
	return fmt.Sprintf("\x1b[;1;1m%v\x1b[0m", s)
}
