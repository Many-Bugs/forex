package starter

type Redis struct {
	Mode      string
	Frequency int
}

func (m *Redis) Builder() error {
	return nil
}
