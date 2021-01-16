package model

import (
	"git.code.oa.com/leoshli/code-generator/conf"
	"git.code.oa.com/leoshli/code-generator/utils/goenv"
	"git.code.oa.com/leoshli/code-generator/utils/nameutil"
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
	}

	result.Common = Common{
		ObjectMark: ObjectMark,
		ServerMark: ServerMark,
	}

	return result, nil
}

func genDir(metadata *Metadata) string {
	return metadata.ModulePath
}
