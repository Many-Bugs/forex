package starter

type Mysql struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
	IsWebService bool
	Error        error
	ModelAddrs   []interface{}
}

func (m *Mysql) Builder(c *Content) error {
	return nil
}

func (m *Mysql) AutoMigrateAddr(obj interface{}) {
	m.ModelAddrs = append(m.ModelAddrs, obj)
	return
}

func (m *Mysql) New() {
	m.ModelAddrs = make([]interface{}, 0)
	return
}

func (m *Mysql) createDatabase() {

}
