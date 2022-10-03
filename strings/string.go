package strings

import (
	"crypto/md5"
	crand "crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/shootz-developer/gtool/constant"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"strings"
	"time"
	"unsafe"

	"github.com/antlabs/strsim"
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

// TrimSpaceSlice 将src中元素逐个调用strings.TrimSpace。
func TrimSpaceSlice(src []string) []string {
	dest := make([]string, 0, len(src))
	for _, s := range src {
		dest = append(dest, strings.TrimSpace(s))
	}
	return dest
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

// IsStringSliceEqual 返回两个字符串列表是否相等。都为nil，返回true。其中之一为nil，返回false。
// 都不是nil，则长度内容顺序都一致才返回true。
func IsStringSliceEqual(lh, rh []string) bool {
	if lh == nil && rh == nil {
		return true
	}

	if lh == nil || rh == nil {
		return false
	}

	if len(lh) != len(rh) {
		return false
	}

	for k, v := range lh {
		if v != rh[k] {
			return false
		}
	}
	return true
}

// SplitSlice 将slice按长度batchSize分成多段
// batchHandler 返回false，表示退出执行
func SplitSlice(slice interface{}, batchSize int, batchHandler func(batch interface{}) bool) {
	if batchHandler == nil {
		return
	}

	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		panic("argument not a slice")
	}

	blocks := int(math.Ceil(float64(rv.Len()) / float64(batchSize)))
	for i := 0; i < blocks; i++ {
		begin := i * batchSize
		end := begin + batchSize
		if end > rv.Len() {
			end = rv.Len()
		}

		batch := rv.Slice(begin, end)
		isContinue := batchHandler(batch.Interface())
		if !isContinue {
			break
		}
	}
}

// BytesToString copy-free的[]byte转string，但注意使用场景限制，不可滥用。
func BytesToString(bytes []byte) string {
	var s string
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	stringHeader.Data = sliceHeader.Data
	stringHeader.Len = sliceHeader.Len
	return s
}

// NullStr 判断是否是空字符串
func NullStr(str string) bool {
	return len([]rune(str)) == 0
}

// NullStrWithDefault 空字符串赋予默认值
func NullStrWithDefault(str string, defaultStr string) (retStr string) {
	if len([]rune(str)) == 0 {
		return defaultStr
	}
	return str
}

// Similarity 计算两个字符串的相似度
func Similarity(s1 string, s2 string) float64 {
	sim := strsim.Compare(s1, s2)
	return sim
}

// StringFilter 字符串过滤，将1,2,3,4,5 -> '1','2','3','4','5'
func StringFilter(ids string) string {
	splitString := strings.Split(ids, constant.Comma)
	videos := fmt.Sprintf("'%s'", splitString[0])
	for i := 1; i < len(splitString); i++ {
		videos = fmt.Sprintf("%s,'%s'", videos, splitString[i])
	}
	return videos
}
