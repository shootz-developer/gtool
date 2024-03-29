package constant

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

	StartTime = "00:00:00"
	EndTime   = "23:59:59"
)

const (
	MySQLDrive = "mysql" // mysql驱动

	SuccessString = "success"

	StandardTime = "2006-01-02 15:04:05" // StandardTime 标准时间

	DefaultPage  = 0  // 默认的页
	DefaultLimit = 10 // 默认的页大小

	MethodPost = "POST"
	MethodGet  = "GET"

	RetryTimes = 3
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

var WeekDayMap = map[string]string{
	"Monday":    "周一",
	"Tuesday":   "周二",
	"Wednesday": "周三",
	"Thursday":  "周四",
	"Friday":    "周五",
	"Saturday":  "周六",
	"Sunday":    "周日",
}

const (
	DayVIP     = "日会员"
	MonthVIP   = "月会员"
	QuarterVIP = "季会员"
	YearVIP    = "年会员"

	DayAllVIP     = "日大会员"
	MonthAllVIP   = "月大会员"
	QuarterAllVIP = "季大会员"
	YearAllVIP    = "年大会员"

	AllIncome    = "所有收入"
	AllVIPIncome = "大会员收入"
)

const (
	Consumer = "Consumer" // C端的类型
	Business = "Business" // B端的类型
)

const (
	DefaultMembership       = 0
	DefaultMergeLimitMember = 2
	DefaultMaxMerge         = 6
	DefaultDownloadLimit    = 20
)
