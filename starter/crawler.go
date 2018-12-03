package starter

type Crawler struct {
	Mode      string
	Frequency int
}

func (m *Crawler) Builder(c *Content) error {
	return nil
}
