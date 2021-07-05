package model

import "git.code.oa.com/leoshli/code-generator/conf"

//Metadata ...
type Metadata struct {
	Name       Name       `json:"name"`
	ModuleName string     `json:"moduleName"`
	GenDir     string     `json:"genDir"`
	ModulePath string     `json:"modulePath"`
	Ctrl       Controller `json:"ctrl"`
	Common     Common     `json:"common"`
}

//Name ...
type Name struct {
	Name              string `json:"name"`              // name from cmd
	Snake             string `json:"snake"`             // eg: resource_manager
	ExportedCamel     string `json:"exportedCamel"`     // ResourceManager
	ExportedCamelList string `json:"exportedCamelList"` // ResourceMangerList
	ShortName         string `json:"shortName"`         // r or m
	UnexportedCamel   string `json:"unexportedCamel"`   // eg: resourceManager
}

//Controller ...
type Controller struct {
	DaoMetricsFlag bool             `json:"daoMetricsFlag"`
	APIServer      bool             `json:"apiServer"`
	Storage        conf.StorageType `json:"storage"`
	TemplateDir    string           `json:"templateDir"`
	ProtoOnly      bool             `json:"protoOnly"`
	AppName        string           `json:"appName"`
	ServerName     string           `json:"serverName"`
}

//Common ...
type Common struct {
	ObjectMark string `json:"objectMark"`
	ServerMark string `json:"serverMark"`
}

// const ...
const (
	ObjectMark = "object@code-generator"
	ServerMark = "service@code-generator"
)

//NewMetadata ...
func NewMetadata() *Metadata {
	return &Metadata{}
}
