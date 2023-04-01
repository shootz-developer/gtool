package utils

import (
	"github.com/shootz-developer/gtool/constant"
	"log"
	"math"
	"strconv"
)

// EarthDistance 计算距离
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = lat1 * math.Pi / 180.0
	lng1 = lng1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0
	lng2 = lng2 * math.Pi / 180.0
	a := lat1 - lat2
	b := lng1 - lng2
	s := 2 * math.Asin(
		math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(b/2), 2)))
	s = s * 6378.137
	s = math.Round(s*10000) / 10000
	return s
}

// GetPage 获取页
func GetPage(pageStr string) int {
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Printf("Page string to int err: [%+v]", err)
		return constant.DefaultPage
	}

	return page
}

// GetLimit 获取页大小
func GetLimit(pageStr string) int {
	limit, err := strconv.Atoi(pageStr)
	if err != nil {
		log.Printf("Page string to int err: [%+v]", err)
		return constant.DefaultLimit
	}

	return limit
}
