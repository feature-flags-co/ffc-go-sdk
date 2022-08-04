package utils

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var alphabetsMap map[string]string = map[string]string{
	"0": "Q",
	"1": "B",
	"2": "W",
	"3": "S",
	"4": "P",
	"5": "H",
	"6": "D",
	"7": "X",
	"8": "Z",
	"9": "U",
}

func BuildToken(envSecret string) string {
	text := strings.TrimRight(envSecret, "=")
	now := time.Now().UnixNano() / 1e6
	timestampCode := encodeNumber(now, len(strconv.FormatInt(int64(now), 10)))
	start := math.Max(math.Floor(rand.Float64()*float64(len(text))), 2)

	part1 := encodeNumber(int64(start), 3)
	part2 := encodeNumber(int64(len(timestampCode)), 2)
	part3 := substring(text, 0, int(start))
	part4 := timestampCode
	part5 := substring(text, int(start), len(text))
	result := fmt.Sprintf("%s%s%s%s%s", part1, part2, part3, part4, part5)
	return result
}

func encodeNumber(number int64, length int) string {
	str := "000000000000" + strconv.FormatInt(number, 10)
	numberWithLeadingZeros := substring(str, len(str)-length, len(str))
	strList := strings.Split(numberWithLeadingZeros, "")
	var encodeStr string
	for _, v := range strList {
		encodeStr = encodeStr + alphabetsMap[v]
	}
	return encodeStr

}

func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	var substring = ""
	for i := start; i < end; i++ {
		substring += string(r[i])
	}

	return substring
}

// DefaultHeaders set dafault header for ffc request
func DefaultHeaders(envSecret string) map[string]string {
	headers := make(map[string]string, 0)
	headers["envSecret"] = envSecret
	headers["User-Agent"] = "ffc-go-server-sdk4"
	headers["Content-Type"] = "application/json"
	return headers
}

// HeaderBuilderFor convert http config header to request headers
// @Param httpConfig
// @Return http headers
