package model

import (
	"git.code.oa.com/tencent_abtest/code-generator/conf"
	"git.code.oa.com/tencent_abtest/code-generator/utils/goenv"
	"git.code.oa.com/tencent_abtest/code-generator/utils/nameutil"
	"github.com/pkg/errors"
)

// GenMetadata gen metadata from env and cmd args
func GenMetadata() (*Metadata, error) {

	result := NewMetadata()

	// get module-name
	moduleName, err := goenv.GetModuleName()
	if err != nil {
		return nil, errors.Wrapf(err, "GetModuleName fail")
	}
	result.ModuleName = moduleName

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

	// gen gen dir
	modulePath, err := goenv.GetModulePath()
	if err != nil {
		return nil, errors.Wrapf(err, "GetModulePath fail")
	}
	result.ModulePath = modulePath
	result.GenDir = genDir(result)

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

	result.Common = Common{
		ObjectMark: ObjectMark,
		ServerMark: ServerMark,
	}

	err = validateResult(result)
	if err != nil {
		return nil, errors.Wrapf(err, "validate metadata")
	}

	return result, nil
}

func genDir(metadata *Metadata) string {
	return metadata.ModulePath
}

func validateResult(metadata *Metadata) error {

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
