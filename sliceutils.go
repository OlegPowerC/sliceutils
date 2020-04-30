package sliceutils

import "fmt"

func SliceFromArrayByte(start int, end int, source []byte, debuglevel int) (err error, retarr []byte) {
	if debuglevel > 0 {
		fmt.Println("start", start, "end", end, "len", len(source))
	}

	if start > len(source) || end > len(source) {
		return fmt.Errorf("%d", 1), nil

	}

	if start >= end {
		return fmt.Errorf("%d", 2), nil
	} else {
		if end < len(source) {
			return nil, source[start:end]
		} else {
			return nil, source[start:]
		}

	}
}

func SliceFromArrayString(start int, end int, source string, debuglevel int) (err error, retarr string) {
	if debuglevel > 0 {
		fmt.Println("start", start, "end", end, "len", len(source))
	}

	if start > len(source) || end > len(source) {
		return fmt.Errorf("%d", 3), ""
	}

	if start >= end {
		return fmt.Errorf("%d", 4), ""
	} else {
		if end < len(source) {
			return nil, source[start:end]
		} else {
			return nil, source[start:]
		}

	}
}
