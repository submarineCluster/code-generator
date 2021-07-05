package conf

// var ...
var (
	ResourceName string
	Verbose      bool
	DaoMetrics   bool
	APIServer    bool
	StorageT     string
	TemplateDir  string
	ProtoOnly    bool
	AppName      string
	ServerName   string
	CacheEnable  bool
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
