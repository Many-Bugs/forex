package starter

import (
	"fmt"
	"forex/systems"
	"log"
	"os"
	"time"
)

type LogHierarchyOrder int

const (
	ALL LogHierarchyOrder = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

type Logger struct {
	LoggerInstance
	// log setting
	IsLog            bool
	LogDepth         int
	RootPath         string
	LogSaveName      string
	LogSavePath      string
	LogFileExtension string

	// HTTP traffic log
	IsHTTPMessageLog            bool
	HTTPMessageLogSavePath      string
	HTTPMessageLogSaveName      string
	HTTPMessageLogFileExtension string
	DefaultLoggingOrderTag      []string

	Prefix string

	// server setting
	IsServerLog            bool
	ServerLogSavePath      string
	ServerLogSaveName      string
	ServerLogFileExtension string

	// web panel
	IsTerminal     bool
	IsWebPanel     bool
	IsSaveToMongo  bool
	CollectionName string
	Host           string
	Port           int
	Domain         string
}

type LoggerInstance struct {
	File             *os.File
	ServerFile       *os.File
	HTTPMessagesFile *os.File
	Logger           *log.Logger
}

func (m *Logger) Builder(c *Content) error {

	var err error

	m.RootPath = systems.ReplaceSplit(c.App.RootPath)

	if m.IsLog {
		m.LogSaveName = systems.ReplaceSplit(m.LogSaveName)
		m.LogSavePath = systems.ReplaceSplit(m.LogSavePath)

		m.File, err = systems.MustOpen(
			fmt.Sprintf("%s%s.%s", m.LogSaveName, time.Now().Local().Format("2006-01-02 15:04:05"), m.LogFileExtension),
			fmt.Sprintf("%s%s", m.RootPath, m.LogSavePath),
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if m.IsServerLog {
		m.ServerLogSavePath = systems.ReplaceSplit(m.ServerLogSavePath)
		m.ServerLogSaveName = systems.ReplaceSplit(m.ServerLogSaveName)

		m.ServerFile, err = systems.MustOpen(
			fmt.Sprintf("%s%s.%s", m.ServerLogSaveName, time.Now().Local().Format("2006-01-02 15:04:05"), m.ServerLogFileExtension),
			fmt.Sprintf("%s%s", m.RootPath, m.ServerLogSavePath),
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if m.IsHTTPMessageLog {
		m.HTTPMessageLogSavePath = systems.ReplaceSplit(m.HTTPMessageLogSavePath)
		m.HTTPMessageLogSaveName = systems.ReplaceSplit(m.HTTPMessageLogSaveName)

		m.HTTPMessagesFile, err = systems.MustOpen(
			fmt.Sprintf("%s%s.%s", m.HTTPMessageLogSaveName, time.Now().Local().Format("2006-01-02 15:04:05"), m.HTTPMessageLogFileExtension),
			fmt.Sprintf("%s%s", m.RootPath, m.HTTPMessageLogSavePath),
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if m.IsLog || m.IsServerLog || m.IsHTTPMessageLog {
		m.Logger = log.New(m.File, m.Prefix, log.LstdFlags)
	}

	return nil
}

func (m *Logger) Starter(c *Content) error {
	return nil
}

func (m *Logger) Router(s *Server) {
	
}

