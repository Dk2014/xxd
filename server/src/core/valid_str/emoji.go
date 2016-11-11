package valid_str

import (
	"errors"
	"strings"
)

func StripEmoji(src string) (err error) {
	runes := []rune(src)

	if strings.ContainsAny(src, " .<>&'\"_@#") {
		return errors.New("could not contains ' .<>&'\"_@#' ")
	}

	for _, v := range runes {
		if v >= 0x10000 && v <= 0x10FFFF {
			return errors.New("could not contains emoji")
		}
	}
	return
}
