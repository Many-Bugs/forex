package starter

type Redis struct {
	Mode      string
	Frequency int
}

func (m *Redis) Builder(c *Content) error {
	return nil
}
