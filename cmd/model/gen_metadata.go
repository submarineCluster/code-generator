package model

import (
	"github.com/pkg/errors"
	"github.com/submarineCluster/code-generator/conf"
	"github.com/submarineCluster/code-generator/utils/goenv"
	"github.com/submarineCluster/code-generator/utils/nameutil"
)

// GenMetadata gen metadata from env and cmd args
func GenMetadata() (*Metadata, error) {

	result := NewMetadata()

	// gen model
	modelData := Name{
		Name:              conf.ResourceName,
		Snake:             nameutil.ToSnakeCaseName(conf.ResourceName),
		ExportedCamel:     nameutil.ToExportedCamel(conf.ResourceName),
		ShortName:         nameutil.ToShortName(conf.ResourceName),
		UnexportedCamel:   nameutil.ToUnexportedCamel(conf.ResourceName),
		ExportedCamelList: nameutil.ToExportedCamel(conf.ResourceName) + "List",
	}
	result.Name = modelData

	// controller
	result.Ctrl = Controller{
		DaoMetricsFlag: conf.DaoMetrics,
		APIServer:      conf.APIServer,
		Storage:        conf.StorageType(conf.StorageT),
		TemplateDir:    conf.TemplateDir,
		ProtoOnly:      conf.ProtoOnly,
		AppName:        conf.AppName,
		ServerName:     conf.ServerName,
		CacheEnable:    conf.CacheEnable,
	}

	// common
	result.Common = Common{
		ObjectMark: ObjectMark,
		ServerMark: ServerMark,
	}
	// modulePath
	modulePath, err := goenv.GetModulePath()
	if err != nil {
		return nil, errors.Wrapf(err, "GetModulePath fail")
	}
	result.ModulePath = modulePath
	result.GenDir = genDir(result)

	if !conf.ProtoOnly { // 可以在项目外生成proto文件
		// gen gen dir
		// get module-name
		moduleName, err := goenv.GetModuleName()
		if err != nil {
			return nil, errors.Wrapf(err, "GetModuleName fail")
		}
		result.ModuleName = moduleName
	}
	err = validateResult(result)
	if err != nil {
		return nil, errors.Wrapf(err, "validate metadata")
	}

	return result, nil
}

//genDir ...
func genDir(metadata *Metadata) string {
	return metadata.ModulePath
}

//validateResult ...
func validateResult(metadata *Metadata) error {

	// protoOnly
	if metadata.Ctrl.ProtoOnly {
		if len(metadata.Ctrl.ServerName) == 0 {
			return errors.Errorf("required serverName")
		}
	}
	if metadata.Ctrl.TemplateDir == "ddd" || metadata.Ctrl.TemplateDir == "trpc" {
		if len(metadata.Ctrl.ServerName) == 0 {
			return errors.Errorf("required serverName")
		}
	}

	// serverName
	if len(metadata.Ctrl.ServerName) > 0 {
		for i, s := range metadata.Ctrl.ServerName {
			if i == 0 {
				if s < 'a' || s > 'z' {
					return errors.Errorf("invalid serverName, required snakeCase")
				}
			}

			if i == len(metadata.Ctrl.ServerName)-1 {
				if s == '_' {
					return errors.Errorf("invalid serverName, required snakeCase")
				}
			}
			if 'a' <= s && s <= 'z' || '0' < s && s <= '9' || s == '_' {
				continue
			}
			return errors.Errorf("invalid serverName, required snakeCase")
		}
	}
	return nil
}
