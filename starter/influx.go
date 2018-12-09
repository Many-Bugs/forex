package starter

import (
	"forex/utils"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

type Influx struct {
	InfluxInstance
	Username                       string
	Password                       string
	Host                           string
	Port                           int
	DatabaseName                   string
	IsWebService                   bool
	Error                          error
	MaximumIdleConnection          int
	MaximumOpenConnection          int
	MaximumConnectionRetry         int
	MinimumRetryDuration           int
	MaximumConnectionKeepAliveTime int
}

type InfluxInstance struct {
	AliveConnectionCount int
	Client               client.Client
	Response             *client.Response
}

func (m *Influx) Builder(c *Content) error {
	m.CreateDatabase()
	return nil
}

func (m *Influx) CreateDatabase() {
	if close := m.Connector(); close != nil {
		defer close(m)
		m.Response, m.Error =
			m.Client.Query(client.NewQuery("CREATE DATABASE "+m.DatabaseName, "", ""))
		utils.AssertErr(m.Error)
		return
	}
}

func (m *Influx) Connector() func(*Influx) error {
	close := func(m *Influx) error {
		m.AliveConnectionCount--
		return m.Client.Close()
	}
	m.recursionCall(
		func() error {
			m.Client, m.Error = client.NewHTTPClient(client.HTTPConfig{
				Addr:     "http://" + m.Host + ":" + strconv.Itoa(m.Port),
				Username: m.Username,
				Password: m.Password,
			})
			return m.Error
		},
		m.MaximumConnectionRetry,
		m.MinimumRetryDuration,
		false)
	if !utils.AssertErr(m.Error) {
		m.AliveConnectionCount++
		return close
	}
	return nil
}

func (m *Influx) recursionCall(f func() error, count, duration int, done bool) bool {
	if !done {
		m.Error = f()
		count--
	}
	if count > 0 && m.Error == nil {
		return true
	} else if count == 0 && m.Error != nil {
		return true
	} else {
		time.Sleep(time.Duration(duration) * time.Second)
	}
	return m.recursionCall(f, count, duration, false)
}
