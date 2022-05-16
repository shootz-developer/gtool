package constant

const (
	Success          = 0 // 请求成功的时候的返回值
	Fail             = 1 // 请求失败的时候的返回值
	DataEmpty        = 2 // 请求的数据为空时候的返回值
	ParameterMiss    = 3 // 请求所需的参数部分缺失
	ParameterInvalid = 4 // 请求的参数存在不合法的情况
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
)

const (
	MySQLDrive = "mysql" // mysql驱动
)
