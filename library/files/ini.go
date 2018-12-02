package files

import (
	"fmt"

	"github.com/go-ini/ini"
)

type INIFileBinding struct {
	File string
	Dir  string
}

func (i INIFileBinding) BindFile(obj interface{}) error {
	f, err := ini.Load(i.Dir)
	if err != nil {
		return fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return mapTo(f, "", obj)
}

func mapTo(f *ini.File, section string, v interface{}) error {
	return f.MapTo(v)
}
