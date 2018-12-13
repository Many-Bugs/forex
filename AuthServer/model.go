package AuthServer

import (
	"forex/starter"
)

type User struct {
	starter.MysqlModel
	Username string
	ParentID int
	Level    int
}

func ModelAddrs() []interface{} {
	return []interface{}{
		&User{},
	}
}
