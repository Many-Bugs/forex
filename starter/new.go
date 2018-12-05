package starter

import (
	"forex/debugs"
	"forex/library/files"
	"forex/systems"
)

const (
	INISetting  = "setting.ini"
	JSONSetting = "setting.json"
)

var (
	configFiles = []string{
		INISetting,
		JSONSetting,
	}
)

type Content struct {
	ConfigFile string
	App        App
	Logger     Logger
	Server     Server
	Mysql      Mysql
	Mongo      Mongo
	Influx     Influx
	Redis      Redis
	Crawler    Crawler
}

func Default() *Content {

	content := &Content{
		ConfigFile: INISetting,
	}
	for i := 0; i < len(configFiles); i++ {
		if !systems.CheckNotExist(configFiles[i]) {
			content.ConfigFile = configFiles[i]
		}
	}

	files.BindFileToObj(content.ConfigFile, content)

	content.Builder(&content.App)
	content.Builder(&content.Logger)
	content.Builder(&content.Server)
	content.Builder(&content.Mysql)
	content.Builder(&content.Mongo)
	content.Builder(&content.Influx)
	content.Builder(&content.Redis)
	content.Builder(&content.Crawler)

	debugs.PrintStructureWithField(content)

	return content
}

func (m *Content) Builder(b Builder) error {
	b.Builder(m)
	return nil
}

func (m *Content) Starter(s Starter) error {
	return nil
}
