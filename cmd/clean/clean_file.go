package clean

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"

	"github.com/submarineCluster/code-generator/cmd/model"
	"github.com/submarineCluster/code-generator/utils/log"
)

//clean ...
func clean(metadata *model.Metadata) error {

	for _, path := range genResourceDir(metadata) {
		err := removeAll(path)
		if err != nil {
			log.Printf("removeAll %v fail: %v", path, err)
			continue
		}
	}

	return nil
}

//removeAll ...
func removeAll(name string) error {
	// remove file -rf
	err := os.RemoveAll(name)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("os RemoveAll %v fail: %v", name, err)
	}
	dir, err := ioutil.ReadDir(name)
	if err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "ioutil ReadDir fail, name=%v", name)
	}
	for _, d := range dir {
		path := path.Join([]string{name, d.Name()}...)
		err = os.RemoveAll(path)
		if err != nil {
			log.Printf("os RemoveAll %v fail: %v", path, err)
			continue
		}
	}
	return nil
}

//genResourceDir ...
func genResourceDir(metadata *model.Metadata) []string {
	return []string{
		metadata.GenDir + "/" + metadata.Name.Snake,
		metadata.GenDir + "/api/" + metadata.Name.Snake,
	}
}
