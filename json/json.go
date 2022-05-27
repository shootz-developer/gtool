package json

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/prometheus/common/log"
)

// String2Json 字符串转换为json
func String2Json(j string, v interface{}) error {
	return json.Unmarshal([]byte(j), v)
}

// Map2JsonString map对象转化成json string
func Map2JsonString(data map[string]interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

// IsJSON 判断字符串是否是合法的JSON字符串
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// JSONToMap 将JSON字符串转化为map——value是任意类型
func JSONToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		log.Infof("JSON解析异常:%+v", err)
		return nil, err
	}
	return m, nil
}

// GetJSONKeyValue 解码未知结构的json并获取指定键值
func GetJSONKeyValue(unknownJSON interface{}, key string) (value interface{}) {
	unknownJSONMap, ok := unknownJSON.(map[string]interface{})
	if ok {
		return unknownJSONMap[key]
	}
	return
}

//GetJSONKeyStringValue 获取JSON对象中指定key的字符串值
func GetJSONKeyStringValue(object interface{}, key string) string {
	val := GetJSONKeyValue(object, key)
	strVal := fmt.Sprintf("%v", val)
	return strVal
}

// DumpString 解析为JSON字符串
func DumpString(v interface{}) (str string) {
	bs, err := json.Marshal(v)
	b := bytes.Buffer{}
	if err != nil {
		b.WriteString("{err:\"JSON格式错误.")
		b.WriteString(err.Error())
		b.WriteString("\"}")
	} else {
		b.Write(bs)
	}
	str = b.String()
	return str
}

// PrintUnknownJSON 解码未知结构的json
func PrintUnknownJSON(unknownJSON interface{}) {
	unknownJSONMap, ok := unknownJSON.(map[string]interface{})
	if ok {
		for k, v := range unknownJSONMap {
			switch kv := v.(type) {
			case string:
				fmt.Println(k, "is string", kv)
			case int:
				fmt.Println(k, "is int", kv)
			case bool:
				fmt.Println(k, "is bool", kv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range kv {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "类型未知")
			}
		}
	}
}
