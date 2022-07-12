package filter

import (
	"log"
	"os"
	"strconv"
)

// Size filter files that are smaller than the specified size.
func Size(path string, args []string) bool {
	if len(args) != 1 {
		log.Println("Size: requires an argument")
	}

	minSize, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		log.Println("Size: must be a number")
	}

	st, err := os.Stat(path)
	if err != nil {
		return false
	}

	size := st.Size()

	if size >= minSize {
		return true
	}

	return false
}

// Consumes filter files that are not in the specified size range.
func Consumes(path string, args []string) bool {
	log.Println("Consumes: NOT IMPLEMENTED!")
	return false
}
