package files

import (
	"path"
)

type BindingFiles interface {
	BindFile(interface{}) error
}

func GetFileType(dir string) BindingFiles {
	dir = ReplaceSplit(dir)
	switch path.Ext(dir) {
	case ".json":
		return JSONFileBinding{Dir: dir}
	case ".ini":
		return INIFileBinding{Dir: dir}
	default:
		return JSONFileBinding{Dir: dir}
	}
}

// Api: the file name and the pointer of the struct variable
func BindFileToObj(path string, obj interface{}) error {
	b := GetFileType(path)
	return b.BindFile(obj)
}
