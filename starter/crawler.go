package starter

type Crawler struct {
	Mode      string
	Frequency int
}

func (m *Crawler) Builder(c *Content) error {
	return nil
}

func (m *Crawler) Starter(c *Content) error {
	return nil
}

func (m *Crawler) Router(s *Server) {

}
