package conf

// var ...
var (
	ResourceName string // 资源名
	Verbose      bool   // 是否开启日志输出
	DaoMetrics   bool   // 是否上报dao metrics
	APIServer    bool   // 是否生成 server 层代码
	StorageT     string // 存储层类型
	TemplateDir  string // 模板文件目录
	ProtoOnly    bool   // 只生成协议
	AppName      string // app name
	ServerName   string // sever name
	CacheEnable  bool   // 是否开启缓存
)

//StorageType ...
type StorageType string

// const ...
const (
	StorageTypeMySQL StorageType = "mysql"
	StorageTypeMongo StorageType = "mongo"
)

//Inst ...
type Inst struct {
	GoPkg string `json:"goPkg"`
}

// const ...
const (
	SufferString = ".temp"
)
