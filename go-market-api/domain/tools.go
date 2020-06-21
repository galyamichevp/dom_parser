package domain

import (
	"regexp"
	"strconv"
	"strings"
)

// FindPercentValue - find percent value in string
func FindPercentValue(source string) (float64, error) {
	r, _ := regexp.Compile("[0-9,.]+")

	res := r.FindString(source)

	res = strings.ReplaceAll(res, ",", "")

	f, err := strconv.ParseFloat(res, 64)
	if err == nil {
		return f, nil
	}

	return -1, err
}
