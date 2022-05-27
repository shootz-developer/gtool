package collection

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

const (
	keyNotFoundErr = "在map中key不存在"
	notStringErr   = "不为string类型"
	notFloat64Err  = "不为float64类型"
	notInt64Err    = "不为int64类型"
	notIntErr      = "不为int类型"
)

// InsertStringMap 插入map的简单封装，主要是判断map是否是空的，如果为空就构造一个map
func InsertStringMap(srcMap map[string]string, key, value string) map[string]string {
	if srcMap == nil {
		srcMap = map[string]string{
			key: value,
		}
		return srcMap
	}
	srcMap[key] = value
	return srcMap
}

// StructToMap struct转为map 因为struct是不能遍历的
func StructToMap(stu interface{}) map[string]interface{} {
	t := reflect.TypeOf(stu)
	v := reflect.ValueOf(stu)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

// MapToStruct 强类型map转换struct
func MapToStruct(input interface{}, output interface{}, tags ...string) error {
	tag := ""
	if len(tags) >= 1 {
		tag = tags[0]
	}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: tag,
		Result:  output,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

// GetStringFromMap 从map中安全获取string值
func GetStringFromMap(notSafeMap map[string]interface{}, key string) (string, error) {
	var ok bool
	if _, ok := notSafeMap[key]; !ok {
		return "", fmt.Errorf(key + keyNotFoundErr)
	}
	var v string
	if v, ok = notSafeMap[key].(string); !ok {
		return "", fmt.Errorf(key + notStringErr)
	}
	return v, nil
}

// GetFloat64FromMap 从map中安全获取float64值
func GetFloat64FromMap(notSafeMap map[string]interface{}, key string) (float64, error) {
	var ok bool
	if _, ok := notSafeMap[key]; !ok {
		return float64(0), fmt.Errorf(key + keyNotFoundErr)
	}
	var v float64
	if v, ok = notSafeMap[key].(float64); !ok {
		return float64(0), fmt.Errorf(key + notFloat64Err)
	}
	return v, nil
}

// GetInt64FromMap 从map中安全获取int64值
func GetInt64FromMap(notSafeMap map[string]interface{}, key string) (int64, error) {
	var ok bool
	if _, ok := notSafeMap[key]; !ok {
		return int64(0), fmt.Errorf(key + keyNotFoundErr)
	}
	var v int64
	if v, ok = notSafeMap[key].(int64); !ok {
		return int64(0), fmt.Errorf(key + notInt64Err)
	}
	return v, nil
}

// GetIntFromMap 从map中安全获取int值
func GetIntFromMap(notSafeMap map[string]interface{}, key string) (int, error) {
	var ok bool
	if _, ok := notSafeMap[key]; !ok {
		return 0, fmt.Errorf(key + keyNotFoundErr)
	}
	var v int
	if v, ok = notSafeMap[key].(int); !ok {
		return 0, fmt.Errorf(key + notIntErr)
	}
	return v, nil
}
