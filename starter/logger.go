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
	// log setting
	LogDepth int
	RootPath string

	LogSaveName      string
	LogSavePath      string
	LogFileExtension string

	HTTPMessageLogSavePath      string
	HTTPMessageLogSaveName      string
	HTTPMessageLogFileExtension string

	DefaultLoggingOrderTag []string

	Prefix string

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

	File             *os.File
	ServerFile       *os.File
	HTTPMessagesFile *os.File
	Logger           *log.Logger
}

func (m *Logger) Builder(c *Content) error {

	var err error

	m.RootPath = systems.ReplaceSplit(c.App.RootPath)

	m.LogSaveName = systems.ReplaceSplit(m.LogSaveName)
	m.LogSavePath = systems.ReplaceSplit(m.LogSavePath)

	m.ServerLogSavePath = systems.ReplaceSplit(m.ServerLogSavePath)
	m.ServerLogSaveName = systems.ReplaceSplit(m.ServerLogSaveName)

	m.HTTPMessageLogSavePath = systems.ReplaceSplit(m.HTTPMessageLogSavePath)
	m.HTTPMessageLogSaveName = systems.ReplaceSplit(m.HTTPMessageLogSaveName)

	m.File, err = systems.MustOpen(
		fmt.Sprintf("%s%s.%s", m.LogSaveName, time.Now().Local().Format("2006-01-02 15:04:05"), m.LogFileExtension),
		fmt.Sprintf("%s%s", m.RootPath, m.LogSavePath),
	)
	if err != nil {
		log.Fatalln(err)
	}
	m.ServerFile, err = systems.MustOpen(
		fmt.Sprintf("%s%s.%s", m.ServerLogSaveName, time.Now().Local().Format("2006-01-02 15:04:05"), m.ServerLogFileExtension),
		fmt.Sprintf("%s%s", m.RootPath, m.ServerLogSavePath),
	)
	if err != nil {
		log.Fatalln(err)
	}
	m.HTTPMessagesFile, err = systems.MustOpen(
		fmt.Sprintf("%s%s.%s", m.HTTPMessageLogSaveName, time.Now().Local().Format("2006-01-02 15:04:05"), m.HTTPMessageLogFileExtension),
		fmt.Sprintf("%s%s", m.RootPath, m.HTTPMessageLogSavePath),
	)
	if err != nil {
		log.Fatalln(err)
	}

	m.Logger = log.New(m.File, m.Prefix, log.LstdFlags)

	return nil
}
