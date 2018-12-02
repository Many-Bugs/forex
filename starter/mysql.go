package starter

type Mysql struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

func (m *Mysql) Builder() error {
	return nil
}
