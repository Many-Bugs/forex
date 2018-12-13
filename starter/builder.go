package starter

import (
	"forex/library/files"
	"forex/systems"
)

type Builder interface {
	Builder(*Content) error
}

var (
	_ Builder = &Config{}
	_ Builder = &Logger{}
	_ Builder = &App{}
	_ Builder = &Server{}
	_ Builder = &Mysql{}
	_ Builder = &Mongo{}
	_ Builder = &Influx{}
	_ Builder = &Redis{}
)

func DefaultBuilder(configFiles []string) *Content {
	content := &Content{}
	for i := 0; i < len(configFiles); i++ {
		if !systems.CheckNotExist(configFiles[i]) {
			content.ConfigFile = configFiles[i]
		}
	}
	files.BindFileToObj(content.ConfigFile, content)
	for i := 0; i < len(content.App.InUseService); i++ {
		content.Builder(content.GetFieldOfStructPointer(content.App.InUseService[i]))
	}
	// for index := 0; index < len(content.App.InUseService); index++ {
	// 	fmt.Println(content.GetFieldOfStructPointer(content.App.InUseService[index]))
	// }
	return content
}

func (m *Content) Builder(b Builder) error {
	return b.Builder(m)
}
