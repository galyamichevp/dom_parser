package domain

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ToDate - parse string to date
func ToDate(source string) time.Time {
	t1, _ := time.Parse("01/02/2006", source)

	return t1
}

// ToTime - parse string to time
func ToTime(source string) time.Time {
	t1, _ := time.Parse("15:04:05", source)

	return t1
}

// ToFloat - parse string to float
func ToFloat(source string) float64 {
	source = strings.ReplaceAll(source, ",", "")

	f, err := strconv.ParseFloat(source, 64)
	if err == nil {
		return f
	}
	return -1
}

// FindPercentValue - find percent value in string
func FindPercentValue(source string) (float64, error) {
	r, _ := regexp.Compile("[0-9,.+-]+")

	res := r.FindString(source)

	res = strings.ReplaceAll(res, ",", "")

	f, err := strconv.ParseFloat(res, 64)
	if err == nil {
		return f, nil
	}

	return -1, err
}

// FindPriceValue - find price value in string
func FindPriceValue(source string) (float64, error) {
	r, _ := regexp.Compile("[0-9,.+-]+")

	res := r.FindString(source)

	res = strings.ReplaceAll(res, ",", "")

	f, err := strconv.ParseFloat(res, 64)
	if err == nil {
		return f, nil
	}

	return -1, err
}

// FindHighLowPriceValue - find high and low value in string
func FindHighLowPriceValue(source string) (float64, float64, error) {
	r, _ := regexp.Compile("[0-9,.+-]+")

	res := r.FindAllString(source, -1)

	if len(res) != 2 {
		return -1, -1, errors.New("FindPercentValue parse error")
	}

	var high float64
	var low float64

	highPriceStr := strings.ReplaceAll(res[0], ",", "")
	lowPriceStr := strings.ReplaceAll(res[1], ",", "")

	high, err := strconv.ParseFloat(highPriceStr, 64)
	if err != nil {
		return -1, -1, err
	}

	low, err = strconv.ParseFloat(lowPriceStr, 64)
	if err != nil {
		return -1, -1, err
	}

	return high, low, err
}
