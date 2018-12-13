package starter

type Mongo struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
	IsWebService bool
}

func (m *Mongo) Builder(c *Content) error {
	return nil
}

func (m *Mongo) Starter(c *Content) error {
	return nil
}

func (m *Mongo) Router(s *Server) {
	
}
