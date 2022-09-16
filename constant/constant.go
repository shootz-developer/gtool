package constant

// 返回1xxxx -> 请求成功
const (
	Success          = 10000 // 请求成功的时候的返回值
	DataEmpty        = 10001 // 请求的数据为空时候的返回值
	ParameterMiss    = 10002 // 请求所需的参数部分缺失
	ParameterInvalid = 10003 // 请求的参数存在不合法的情况
	OpenIDEmpty      = 10004 // OpenID 为空
)

// 返回2xxxx -> MySQL错误
const (
	MySQLQueryError  = 20000 // MySQL的查询错误
	MySQLScanError   = 20001 // MySQL在Scan时候的错误
	MySQLInsertError = 20002 // MySQL的插入错误
	MySQLDeleteError = 20003 // MySQL的删除错误
	MySQLUpdateError = 20004 // MySQL的更新错误
	MySQLCountError  = 20004 // MySQL的计数错误
)

// 返回3xxxx -> 其他类型的错误
const (
	StrconvAtoiError = 30000 // 类型转换错误（主要是string转int）
	ParseTimeError   = 30001 // 日期类型的错误
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

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
