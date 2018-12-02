package starter

import (
	"forex/debug"
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
		debug.DebuggingPrint(`[WARNING] Now require Go version ` + strconv.Itoa(int(v)) + ` or later. `)
	}
	debug.DebuggingPrint(`[WARNING] Building an Config instance. `)
	return nil
}
