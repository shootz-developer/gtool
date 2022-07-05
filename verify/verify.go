package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// VerifyEmailFormat 验证邮箱合法性
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyMobileFormat 验证手机号合法性
func VerifyMobileFormat(mobileNumber string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNumber)
}

// SignIDCard 身份证号码处理
func SignIDCard(idcard string) string {
	cp := idcard
	leth := len(cp)
	return cp[0:4] + " **** **** " + cp[leth-4:]
}

// SignMobilePhone 手机号码处理
func SignMobilePhone(mobile string) string {
	cp := mobile
	leth := len(cp)
	return cp[0:3] + " **** " + cp[leth-4:]
}

// GetUniqueId 生成UniqueId方法
func GetUniqueId() string {
	// 当前毫秒时间戳
	timestamp := time.Now().UnixNano() / 1000000
	s1, _ := rand.Int(rand.Reader, big.NewInt(10))
	s2, _ := rand.Int(rand.Reader, big.NewInt(10))

	seqNo := timestamp * (s1.Int64() + 1) * (s2.Int64() + 1)
	uniqueId := fmt.Sprintf("%d", seqNo)
	return uniqueId
}

// VerifyPassword 验证密码合法性
func VerifyPassword(password string) bool {
	pwdLen := len(password)
	digitFlag := 0
	bigLetterFlag := 0
	smallLetterFlag := 0
	specialLetterFlag := 0
	for i := 0; i < pwdLen; i++ {
		if password[i] >= '0' && password[i] <= '9' {
			digitFlag = 1
		} else if password[i] >= 'a' && password[i] <= 'z' {
			smallLetterFlag = 1
		} else if password[i] >= 'A' && password[i] <= 'Z' {
			bigLetterFlag = 1
		} else if password[i] == '!' || password[i] == '~' || password[i] == '@' || password[i] == '#' ||
			password[i] == '_' || password[i] == '%' || password[i] == '^' || password[i] == '*' {
			specialLetterFlag = 1
		} else {
			return false
		}
	}
	flag := digitFlag + smallLetterFlag + specialLetterFlag + bigLetterFlag
	if pwdLen >= 6 && pwdLen <= 15 && flag >= 3 {
		return true
	}

	return false
}

//VerifyTimeDate 判断时间格式是否正确 - xxxx-xx-xx 00:00:00
func VerifyTimeDate(content string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])\s+(20|21|22|23|[0-1]\d):[0-5]\d:[0-5]\d$`)
	return reg.Match([]byte(content))
}

// VerifyPrice 验证价格
func VerifyPrice(content string) bool {
	reg := regexp.MustCompile(`(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)`)
	return reg.Match([]byte(content))
}

// VerifyUrl 验证网络地址
func VerifyUrl(content string) bool {
	reg := regexp.MustCompile(`http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?`)
	return reg.Match([]byte(content))
}

// VerifyBankCode 验证银行卡
func VerifyBankCode(content string) bool {
	reg := regexp.MustCompile(`^\d{16,21}$`)
	return reg.Match([]byte(content))
}

// VerifyInteger 验证整数
func VerifyInteger(content string) bool {
	reg := regexp.MustCompile(`^-?\\d+$`)
	return reg.Match([]byte(content))
}

// VerifyIdcard 身份证号正确性检查
func VerifyIdcard(idCard string) bool {
	idCard = strings.ToUpper(idCard)

	reg := regexp.MustCompile(`^[0-9]{17}[0-9X]$`)
	if reg.MatchString(idCard) == false {
		return false
	}
	return checkIdCardCode(idCard)
}

// 身份证校验码的计算方法：
//  1、将身份证号码前面的17位数分别乘以不同的加权因子，见： weights
//  2、将这17位数字和加权因子相乘的结果相加，得到的结果再除以11，得到余数 m
//  3、余数m作为位置值，在校验码数组 codes 中找到对应的值，就是身份证号码的第18位校验码
func checkIdCardCode(id string) bool {
	var weights []int = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var codes []string = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	var sum int = 0
	for i := 0; i < 17; i++ {
		n, _ := strconv.Atoi(string(id[i]))
		sum += n * weights[i]
	}

	m := sum % 11

	return codes[m] == id[17:]
}
