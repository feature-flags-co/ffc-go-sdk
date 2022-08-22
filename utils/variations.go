package utils

import (
	"crypto/md5"
	"encoding/binary"
	"math"
	"strconv"
	"strings"
)

func IfKeyBelongsPercentage(key string, percentageRange []float64) bool {

	min := percentageRange[0]
	max := percentageRange[1]
	if min == 0 && max == 1 {
		return true
	}
	percentage := PercentageOfKey(key)
	if percentage >= min && percentage < max {
		return true
	} else {
		return false
	}
}

func PercentageOfKey(key string) float64 {

	asciiStr := strconv.QuoteToASCII(key)
	asciiRetStr := strings.ReplaceAll(asciiStr, "\"", "")
	asciiBytes := []byte(asciiRetStr)
	md5Bytes := md5.Sum(asciiBytes)
	retBytes := make([]byte, 0)
	for _, v := range md5Bytes {
		retBytes = append(retBytes, v)
	}
	num := binary.LittleEndian.Uint32(retBytes)
	ret := math.Abs(float64(num) / math.MinInt32)
	return ret
}
