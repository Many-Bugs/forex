package starter

type Crawler struct {
	Mode      string
	Frequency int
}

func (m *Crawler) Builder() error {
	return nil
}
