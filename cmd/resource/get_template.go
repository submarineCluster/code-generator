package resource

import (
	"fmt"
	"os"

	"github.com/markbates/pkger"
	"github.com/pkg/errors"

	"git.code.oa.com/tencent_abtest/code-generator/cmd/model"
)

// const ...
const (
	// InstSuffixPath ...
	InstSuffixPath = "/inst"
	// ConfSuffixPath ...
	ConfSuffixPath = "/config"
	// defaultTemp ...
	defaultTemp = "/ddd"
	// ProtoPath ...
	ProtoPath = "/proto" // 协议文件路径
)

//GetTemplate ...
func GetTemplate(metadata *model.Metadata) (string, error) {

	tempDirPath := getTemplateDir(metadata)

	tempPath := tempDirPath
	//switch metadata.Ctrl.Storage {
	//case conf.StorageTypeMongo:
	//	//tempPath = tempDirPath + "/mongotemplate"
	//	//if metadata.Ctrl.APIServer {
	//	//	tempPath = tempDirPath + "/mongoapitemplate"
	//	//}
	//	tempPath = tempDirPath + defaultTemp
	//case conf.StorageTypeMySQL:
	//	tempPath = tempDirPath + defaultTemp
	//default:
	//	tempPath = tempDirPath + defaultTemp
	//}
	_, err := pkger.Stat(tempPath)
	if os.IsNotExist(err) {
		return "", errors.Wrapf(err, "os stat fail")
	}
	return tempPath, nil
}

//GetInstTemplatePath ...
func GetInstTemplatePath(metadata *model.Metadata) (string, error) {
	tempDirPath := getTemplateDir(metadata)
	tempPath := tempDirPath + InstSuffixPath

	_, err := pkger.Stat(tempPath)
	if os.IsNotExist(err) {
		return "", errors.Wrapf(err, "os stat fail")
	}
	return tempPath, nil
}

//GetConfTemplatePath ...
func GetConfTemplatePath(metadata *model.Metadata) (string, error) {
	tempDirPath := getTemplateDir(metadata)
	tempPath := tempDirPath + ConfSuffixPath

	_, err := os.Stat(tempPath)
	if os.IsNotExist(err) {
		return "", errors.Wrapf(err, "os stat fail")
	}
	return tempPath, nil
}

// getTemplateDir ...
func getTemplateDir(metadata *model.Metadata) string {

	if metadata.Ctrl.ProtoOnly {
		return fmt.Sprintf("git.code.oa.com/tencent_abtest/code-generator:/template" + ProtoPath)
	}
	if len(metadata.Ctrl.TemplateDir) > 0 {
		return fmt.Sprintf("git.code.oa.com/tencent_abtest/code-generator:/template/") + metadata.Ctrl.TemplateDir
	}
	return fmt.Sprintf("git.code.oa.com/tencent_abtest/code-generator:/template" + defaultTemp)
	//frame, err := callstack.CallFrame(0)
	//if err != nil {
	//	return "", errors.Wrapf(err, "CallFrame fail")
	//}
	//
	//return filepath.Dir(frame.File()), nil
}
