package starter

import (
	"forex/debugs"
	"forex/library/files"
	"forex/systems"
	"strconv"
)

type Config struct {
	File   string
	App    App
	Server Server
	Mysql  Mysql
	Mongo  Mongo
}

func (m *Config) Builder() error {
	files.BindFileToObj(m.File, m)
	version, _ := systems.GetMinimumVersion(m.App.MinimumGoVersion)
	if v, _ := systems.GetMinimumVersion(""); v <= version {
		debugs.DebuggingPrint(`[WARNING] Now require Go version ` + strconv.Itoa(int(v)) + ` or later. `)
	}
	debugs.DebuggingPrint(`[WARNING] Building an Config instance. `)
	return nil
}
