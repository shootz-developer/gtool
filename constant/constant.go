package constant

const (
	Success          = 10000 // 请求成功的时候的返回值
	Fail             = 10001 // 请求失败的时候的返回值
	DataEmpty        = 10002 // 请求的数据为空时候的返回值
	ParameterMiss    = 10003 // 请求所需的参数部分缺失
	ParameterInvalid = 10004 // 请求的参数存在不合法的情况
	RedisError       = 10005 // Redis出现错误
	MySQLError       = 10006 // MySQL出现错误
	ESError          = 10007 // ES出现错误
)

const (
	// SepFlag 竖线分割符
	SepFlag = "|"
	// Comma 逗号
	Comma = ","
	// Hyphen 连字符
	Hyphen = "-"
	// Colon 冒号
	Colon = ":"
	// CommConfigPath 通用配置文件路径
	CommConfigPath = "comm.yaml"

	// TrueFlag 表示真的数值
	TrueFlag = 1
	// FalseFlag 表示假的数值
	FalseFlag = 0

	TypeAdd   = 1
	TypeMinus = -1
)

const (
	MySQLDrive = "mysql" // mysql驱动
)
