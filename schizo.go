package main

import (
	"fmt"
	"strings"
)

func main() {
	var counter int
	var b strings.Builder

	input := strings.ToLower("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

	for _, c := range input {
		r := rune(c)
		if r < 65 || r > 122 {
			b.WriteRune(r)
			continue
		}

		counter++
		if counter%2 == 0 {
			upper := r - 32
			b.WriteRune(upper)
		} else {
			b.WriteRune(r)
		}
	}

	fmt.Println(b.String())
}
