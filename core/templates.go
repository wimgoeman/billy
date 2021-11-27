package billy

import (
	"errors"
	"io/fs"
	"os"
	"path"

	"github.com/wimgoeman/billy/core/util"
)

type Template struct {
	Path     string
	FileInfo fs.FileInfo
}

func GetTemplate(name string) (Template, error) {
	template := Template{}
	templateDir := util.GetTemplateDir()
	template.Path = path.Join(templateDir, name)
	var err error
	template.FileInfo, err = os.Stat(template.Path)
	if os.IsNotExist(err) {
		return template, errors.New("Template " + name + " not found. It should exist at " + template.Path + ".")
	}
	return template, err
}
