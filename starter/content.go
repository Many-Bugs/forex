package starter

import (
	"reflect"
)

type Content struct {
	ConfigFile string
	App        App
	Logger     Logger
	Server     Server
	Mysql      Mysql
	Postgres   Postgres
	Mongo      Mongo
	Influx     Influx
	Redis      Redis
}

func (m *Content) GetFieldOfStructPointer(name string) Builder {
	val := reflect.Indirect(reflect.ValueOf(m))
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().Name() == name {
			return val.Field(i).Addr().Interface().(Builder)
		}
	}
	return nil
}
