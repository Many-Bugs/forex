package starter

type Mysql struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
	IsWebService bool
}

func (m *Mysql) Builder() error {
	return nil
}
