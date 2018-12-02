package starter

import "forex/library/files"

type Content struct {
	ConfigFile string
	App        App
	Server     Server
	Mysql      Mysql
	Mongo      Mongo
}

func Default() *Content {
	content := &Content{
		ConfigFile: "setting.ini",
	}
	files.BindFileToObj(content.ConfigFile, content)
	content.Builder(&content.App)
	content.Builder(&content.Server)
	

	return content
}

func (m *Content) Builder(b Builder) error {
	b.Builder()
	return nil
}
