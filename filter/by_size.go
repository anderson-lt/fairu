// Copyright (C) 2022, Anderson Lizarazo Tellez

package filter

import (
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
)

// See function parseSize.
//
// Examples of match:
//  10MB
//  20 KiB
var sizeRegexp = regexp.MustCompile(`^(\d+)\s?([A-Za-z]*)$`)

const (
	kb = 1000
	mb = kb * 1000
	gb = mb * 1000
	tb = gb * 1000

	kib = 1024
	mib = kib * 1024
	gib = mib * 1024
	tib = gib * 1024
)

// Size filter files that are smaller than the specified size.
func Size(path string, args []string) bool {
	if len(args) != 1 {
		log.Println("Size: Requires exactly an argument")
		return false
	}

	minSize, err := parseSize(args[0])
	if err != nil {
		log.Println("Size: Invalid format:", err)
		return false
	}

	st, err := os.Stat(path)
	if err != nil {
		log.Println("Size: Error accessing the file:", err)
		return false
	}

	return st.Size() >= minSize
}

// Consumes filter files that are not in the specified size range.
func Consumes(path string, args []string) bool {
	if len(args) != 2 {
		log.Println("Consumes: requires an argument")
	}

	minSize, err := parseSize(args[0])
	if err != nil {
		log.Println("Consumes: must be a number:", err)
	}

	maxSize, err := parseSize(args[1])
	if err != nil {
		log.Println("Consumes: must be a number:", err)
	}

	st, err := os.Stat(path)
	if err != nil {
		return false
	}

	// minSize <= st.Size() <= maxSize
	return st.Size() >= minSize && st.Size() <= maxSize
}

// parseSize analyze an expression of the 20MB type and makes it bytes.
func parseSize(size string) (int64, error) {
	// If size is 15MB, sizes is equal to:
	//  [15MB 15 MB]
	sizes := sizeRegexp.FindStringSubmatch(size)
	if sizes == nil {
		return 0, errors.New("invalid expression")
	}

	// Convert to int64.
	sizeInBytes, err := strconv.ParseInt(sizes[1], 10, 0)
	if err != nil {
		return 0, err
	}

	if len(sizes) == 3 {
		switch sizes[2] {
		case "B", "":
			// Do not do anything, this is the unity by default.
		case "KB":
			sizeInBytes *= kb
		case "KiB":
			sizeInBytes *= kib
		case "MB":
			sizeInBytes *= mb
		case "MiB":
			sizeInBytes *= mib
		case "GB":
			sizeInBytes *= gb
		case "GiB":
			sizeInBytes *= gib
		case "TB":
			sizeInBytes *= tb
		case "TiB":
			sizeInBytes *= tib
		default:
			return 0, errors.New("Invalid unit " + sizes[2])
		}
	}

	return sizeInBytes, nil
}
