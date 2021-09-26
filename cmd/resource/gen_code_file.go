package resource

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"git.code.oa.com/tencent_abtest/code-generator/cmd/model"

	"github.com/markbates/pkger"
	"github.com/pkg/errors"

	"git.code.oa.com/tencent_abtest/code-generator/conf"
)

//GenFileByMetadata ...
func GenFileByMetadata(metadata *model.Metadata) error {

	var err error
	//err = generatorFile(metadata, GetInstTemplatePath, func(metadata *model.Metadata) (s string, err error) {
	//	return metadata.ModulePath + InstSuffixPath, nil
	//})
	//if err != nil {
	//	return errors.Wrapf(err, "generator %s fail", InstSuffixPath)
	//}

	//err = generatorFile(metadata, GetConfTemplatePath, func(metadata *model.Metadata) (s string, err error) {
	//	return metadata.ModulePath + ConfSuffixPath, nil
	//})
	//if err != nil {
	//	return errors.Wrapf(err, "generator %s fail", ConfSuffixPath)
	//}

	err = generatorFile(metadata, GetTemplate, func(metadata *model.Metadata) (s string, err error) {
		return metadata.GenDir, nil
	})
	if err != nil {
		return errors.Wrapf(err, "generator %s fail", metadata.Name.Name)
	}
	return nil
}

//GeneratorFile ...
func generatorFile(metadata *model.Metadata, templatePathFn func(metadata *model.Metadata) (string, error),
	targetPathFn func(metadata *model.Metadata) (string, error)) error {

	if verifyIsExist(metadata) {
		return errors.Errorf("genDir=%v is exist", metadata.GenDir)
	}

	templatePath, err := templatePathFn(metadata)
	if err != nil {
		return errors.Wrapf(err, "exec templatePathFn fail")
	}
	targetPath, err := targetPathFn(metadata)
	if err != nil {
		return errors.Wrapf(err, "exec targetPathFn fail")
	}

	err = pkger.Walk(templatePath, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		iPath := strings.TrimPrefix(path, templatePath)

		// path name
		temp, err := template.New("").Parse(iPath)
		if err != nil {
			return errors.Wrapf(err, "temp Parse fail")
		}
		var query []byte
		buf := bytes.NewBuffer(query)
		err = temp.Execute(buf, metadata)
		if err != nil {
			return errors.Wrapf(err, "temp Execute fail")
		}

		iPath = buf.String()
		buf.Reset()

		filePath := filepath.Join(targetPath, iPath)
		filePath = strings.TrimSuffix(filePath, conf.SufferString)

		if _, err = os.Stat(filePath); !os.IsNotExist(err) {
			fmt.Printf("skip %v already exist \n", filePath)
			return nil
		}

		// mkdir -p
		err = os.MkdirAll(filepath.Dir(filePath), 0755)
		if err != nil {
			return errors.Wrapf(err, "os MkdirAll fail")
		}

		// create file
		targetFile, err := os.Create(filePath)
		if err != nil {
			return errors.Wrapf(err, "os Create fail")
		}
		defer targetFile.Close()

		//absPath, err := filepath.Abs(path)
		//if err != nil {
		//	return errors.Wrapf(err, "filepath Abs fail")
		//}

		absPath := path

		//tempFile, err := template.New(filepath.Base(absPath)).ParseFiles(absPath)
		tempFile, err := newTemplate(filepath.Base(absPath), absPath)
		if err != nil {
			return errors.Wrapf(err, "temp ParseFiles fail")
		}
		err = tempFile.Execute(buf, metadata)
		if err != nil {
			return errors.Wrapf(err, "temp Execute fail, filePath=%v", filePath)
		}
		_, err = targetFile.Write(buf.Bytes())
		if err != nil {
			return errors.Wrapf(err, "targetFile Write fail, filePath=%v", filePath)
		}

		return nil
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

//newTemplate ...
func newTemplate(name string, filenames ...string) (*template.Template, error) {
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("template: no files named in call to ParseFiles")
	}

	var err error
	t := template.New(name)
	for _, filename := range filenames {
		f, _ := pkger.Open(filename)
		// Now read it.
		sl, _ := ioutil.ReadAll(f)
		// It can now be parsed as a string.
		t, err = t.Parse(string(sl))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
