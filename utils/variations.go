package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"math"
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

	// TODO
	//这里的逻辑就是
	//1. 把 key 用 ASCII 编码得到 byte[]
	//2. MD5 计算 hash 值 得到 byte[]
	//3. 把第二步计算出来的 byte[] 按位转换成  Int32
	//4. 把第三步计算出来的 Int32 / int.最大值 结果取 double 绝对值
	//
	rbytes := md5.Sum([]byte(key))
	number := fromBytes(rbytes[0], rbytes[1], rbytes[2], rbytes[3])
	ret := math.Abs(float64(number) / math.MinInt64)
	return ret
}

func fromBytes(b1 byte, b2 byte, b3 byte, b4 byte) int64 {
	bts := []byte{(b1<<24 | (b2&255)<<16 | (b3&255)<<8 | b4&255)}
	binBuf := bytes.NewReader(bts)
	var x int64
	binary.Read(binBuf, binary.BigEndian, &x)
	return x

}
