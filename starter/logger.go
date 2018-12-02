package starter

type Logger struct {
	isTerminal bool
	isWebPanel bool
	isMongo    bool
	Collection string
	Port       int
	Host       int
}

func (m *Logger) Builder() error {
	return nil
}
