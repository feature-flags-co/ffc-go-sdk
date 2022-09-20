package utils

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"reflect"
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

// DefaultHeaders set default header for ffc request
func DefaultHeaders(envSecret string) map[string]string {
	headers := make(map[string]string, 0)
	headers["envSecret"] = envSecret
	headers["User-Agent"] = "ffc-go-server-sdk"
	headers["Content-Type"] = "application/json"
	return headers
}

// HeaderBuilderFor convert http config header to request headers
// @Param httpConfig
// @Return http headers
func HeaderBuilderFor(headers map[string]string) http.Header {
	header := http.Header{}
	for k, v := range headers {
		header.Add(k, v)
	}
	return header
}

func ToBool(str string) bool {
	if len(str) == 0 {
		return false
	}
	strLow := strings.ToLower(str)
	if strLow == "yes" || strLow == "true" {
		return true
	} else {
		return false
	}
}

func GetString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	default:
		if v != nil {
			return fmt.Sprintf("%v", result)
		}
	}
	return ""
}
func GetInt(v interface{}) int {
	switch result := v.(type) {
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	default:
		if d := GetString(v); d != "" {
			value, _ := strconv.Atoi(d)
			return value
		}
	}
	return 0
}

func GetInt64(v interface{}) int64 {
	switch result := v.(type) {
	case int:
		return int64(result)
	case int32:
		return int64(result)
	case int64:
		return result
	default:

		if d := GetString(v); d != "" {
			value, _ := strconv.ParseInt(d, 10, 64)
			return value
		}
	}
	return 0
}

func GetInt64Slices(v interface{}) []int64 {

	switch result := v.(type) {

	case []int64:
		return []int64(result)
	default:
		fmt.Println("GetInt64Slices: ", result)
		return nil
	}
}

func GetFloat64(v interface{}) float64 {
	switch result := v.(type) {
	case float64:
		return result
	default:
		if d := GetString(v); d != "" {
			value, _ := strconv.ParseFloat(d, 64)
			return value
		}
	}
	return 0
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func GetBool(v interface{}) bool {
	switch result := v.(type) {
	case bool:
		return result
	default:
		if d := GetString(v); d != "" {
			value, _ := strconv.ParseBool(d)
			return value
		}
	}
	return false
}

func GetByteArray(v interface{}) []byte {
	switch result := v.(type) {
	case []byte:
		return result
	case string:
		return []byte(result)
	default:
		return nil
	}
}

func StringsToInterfaces(keys []string) []interface{} {
	result := make([]interface{}, len(keys))
	for i, k := range keys {
		result[i] = k
	}
	return result
}

type StrTo string

func (f *StrTo) Set(v string) {
	if v != "" {
		*f = StrTo(v)
	} else {
		f.Clear()
	}
}

func (f *StrTo) Clear() {
	*f = StrTo("0x1E")
}

func (f StrTo) Exist() bool {
	return string(f) != string("0x1E")
}

func (f StrTo) Bool() (bool, error) {
	return strconv.ParseBool(f.String())
}

func (f StrTo) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

func (f StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

func (f StrTo) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
}

func (f StrTo) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

func (f StrTo) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

func (f StrTo) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

func (f StrTo) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

func (f StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

func (f StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	return uint64(v), err
}

func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func ToInt64(value interface{}) (d int64) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		panic(fmt.Errorf("ToInt64 need numeric not `%T`", value))
	}
	return
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:len(data)]))
}

func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:len(data)])
}

type argString []string

func (a argString) Get(i int, args ...string) (r string) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

type argAny []interface{}

func (a argAny) Get(i int, args ...interface{}) (r interface{}) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

var DefaultTimeLoc = time.Local

func timeParse(dateString, format string) (time.Time, error) {
	tp, err := time.ParseInLocation(format, dateString, DefaultTimeLoc)
	return tp, err
}

func timeFormat(t time.Time, format string) string {
	return t.Format(format)
}

func indirectType(v reflect.Type) reflect.Type {
	switch v.Kind() {
	case reflect.Ptr:
		return indirectType(v.Elem())
	default:
		return v
	}
}
