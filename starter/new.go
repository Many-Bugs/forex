package starter

import (
	"fmt"
	"forex/library/files"
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
		ConfigFile: "setting.ini",
	}
	files.BindFileToObj(content.ConfigFile, content)
	content.Builder(&content.App)
	content.Builder(&content.Logger)
	content.Builder(&content.Server)

	fmt.Printf("%+v\n", content)
	return content
}

func (m *Content) Builder(b Builder) error {
	b.Builder(m)
	return nil
}
