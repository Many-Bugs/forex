package starter

type Logger struct {
	// log setting
	LogDepth         int
	LogSaveName      string
	LogSavePath      string
	LogFileExtension string

	// server setting
	ServerLogSavePath      string
	ServerLogSaveName      string
	ServerLogFileExtension string

	// web panel
	isTerminal     bool
	isWebPanel     bool
	isSaveToMongo  bool
	CollectionName string
	Host           string
	Port           int
	Domain         string
}

func (m *Logger) Builder() error {

	return nil
}
