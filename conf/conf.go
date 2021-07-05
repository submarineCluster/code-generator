package conf

import (
	"git.code.oa.com/leoshli/code-generator/utils/goenv"
)

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

// var ...
var (
	I *Inst
)

// const ...
const (
	SufferString = ".temp"
)

func init() {
	I = &Inst{
		GoPkg: goenv.GetGoPkg(),
	}
}
