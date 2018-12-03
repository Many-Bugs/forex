package starter

type Mongo struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
	IsWebService bool
}

func (m *Mongo) Builder() error {
	return nil
}
