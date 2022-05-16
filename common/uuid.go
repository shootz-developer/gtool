package common

import uuid "github.com/satori/go.uuid"

// MustUUID 创建一个UUID，如果有错误，则抛出panic
func MustUUID() string {
	u := uuid.NewV4()
	return u.String()
}
