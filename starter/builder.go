package starter

import (
	"fmt"
	"forex/library/files"
	"forex/systems"
)

const (
	INISetting  = "setting.ini"
	JSONSetting = "setting.json"
)

type Builder interface {
	Builder(*Content) error
}

var (
	configFiles = []string{
		INISetting,
		JSONSetting,
	}
)

var (
	_ Builder = &Config{}
	_ Builder = &Logger{}
	_ Builder = &App{}
	_ Builder = &Server{}
	_ Builder = &Mysql{}
	_ Builder = &Mongo{}
	_ Builder = &Influx{}
	_ Builder = &Redis{}
	_ Builder = &Crawler{}
)

func DefaultBuilder() *Content {
	content := &Content{
		ConfigFile: INISetting,
	}
	for i := 0; i < len(configFiles); i++ {
		if !systems.CheckNotExist(configFiles[i]) {
			content.ConfigFile = configFiles[i]
		}
	}
	files.BindFileToObj(content.ConfigFile, content)
	for i := 0; i < len(content.App.InUseService); i++ {
		content.Builder(content.GetFieldStructPointer(content.App.InUseService[i]))
	}
	for index := 0; index < len(content.App.InUseService); index++ {
		fmt.Println(content.GetFieldStructPointer(content.App.InUseService[index]))
	}
	return content
}

func (m *Content) Builder(b Builder) error {
	b.Builder(m)
	return nil
}
