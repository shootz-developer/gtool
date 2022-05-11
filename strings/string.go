package strings

import (
	"crypto/md5"
	crand "crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

// NullStringToString nullstring转换成string
func NullStringToString(nullStrings []sql.NullString) []string {
	var strings []string
	for _, nullString := range nullStrings {
		if nullString.Valid {
			strings = append(strings, nullString.String)
		}
	}
	return strings
}

// SplitString 返回以partition分割的字符串数组
func SplitString(str, partition string) []string {
	splitStr := strings.Split(str, partition)
	return splitStr
}

// CaclMD5 返回str内容对应md5值(16进制表示)
func CaclMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// FilterEmptyString 过滤空字符串
func FilterEmptyString(src []string) []string {
	dest := make([]string, 0, len(src))
	for _, s := range src {
		if s == "" {
			continue
		}

		dest = append(dest, s)
	}
	return dest
}

// Base64Encode 对数据进行 base64 编码
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode 对数据进行 base64 解码
func Base64Decode(s string) (string, error) {
	rs, err := base64.StdEncoding.DecodeString(s)
	return string(rs), err
}

// GetRandomString 随机生成给定长度的字符串
func GetRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GetUniqueID 生成UniqueID方法
func GetUniqueID() string {
	// 当前毫秒时间戳
	timestamp := time.Now().UnixNano() / 1000000
	s1, _ := crand.Int(crand.Reader, big.NewInt(10))
	s2, _ := crand.Int(crand.Reader, big.NewInt(10))

	seqNo := timestamp * (s1.Int64() + 1) * (s2.Int64() + 1)
	uniqueID := fmt.Sprintf("%d", seqNo)
	return uniqueID
}
