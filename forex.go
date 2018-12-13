package main

import (
	"forex/AuthServer"
	"forex/starter"
)

const (
	AUTH_SERVER = iota
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

func main() {
	content := starter.DefaultBuilder(configFiles)
	switch content.App.ModuleID {
	case AUTH_SERVER:
		obj := AuthServer.AuthServer{}
		obj.Content = content
		obj.Starter()
	}
}
