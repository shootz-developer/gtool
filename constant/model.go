package constant

// ErrMessage 错误返回的结构体
type ErrMessage struct {
	Err     error
	ErrCode int
}

// MysqlProxy mysql的proxy
type MysqlProxy struct {
	Account  string
	Password string
	Host     string
	Port     string
	DBName   string
}
