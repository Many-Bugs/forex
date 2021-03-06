package starter

import (
	"forex/library/debugs"
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

func (m *Config) Builder(c *Content) error {
	files.BindFileToObj(m.File, m)
	version, _ := systems.GetMinimumVersion(m.App.MinimumGoVersion)
	if v, _ := systems.GetMinimumVersion(""); v <= version {
		debugs.DebuggingPrint(`[WARNING] Now require Go version ` + strconv.Itoa(int(v)) + ` or later. `)
	}
	debugs.DebuggingPrint(`[WARNING] Building an Config instance. `)
	return nil
}

func (m *Config) Starter(c *Content) error {
	return nil
}

func (m *Config) Router(s *Server) {

}
