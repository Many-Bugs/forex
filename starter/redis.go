package starter

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	RedisIntance
	Mode                           string
	Host                           string
	Port                           int
	Password                       string
	Error                          error
	MaximumIdleConnection          int
	MaximumActiveConnection        int
	IdleConnectionTimeout          int
	MaximumConnectionKeepAliveTime int
}

type RedisIntance struct {
	mu        sync.Mutex
	RedisPool *redis.Pool
}

func (m *Redis) Builder(c *Content) error {
	if err := m.setRedisPool(); err != nil {
		return err
	}
	return nil
}

func (m *Redis) Starter(c *Content) error {
	return nil
}

func (m *Redis) setRedisPool() (err error) {
	m.RedisPool = &redis.Pool{
		MaxIdle:         m.MaximumIdleConnection,
		MaxActive:       m.MaximumActiveConnection,
		IdleTimeout:     time.Duration(m.IdleConnectionTimeout),
		MaxConnLifetime: time.Duration(m.MaximumConnectionKeepAliveTime),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", m.Host+":"+strconv.Itoa(m.Port))
			if err != nil {
				return nil, err
			}
			if m.Password != "" {
				if _, err := c.Do("AUTH", m.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	_, err = m.DO("PING")
	return
}

func (m *Redis) getConnection() redis.Conn {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.validate() {
		return m.RedisPool.Get()
	}
	return nil
}

func (m *Redis) recursionCall(f func() error, count, duration int, done bool) bool {
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

func (m *Redis) validate() bool {
	return m.RedisPool != nil &&
		m.RedisPool.ActiveCount() <= m.MaximumActiveConnection &&
		m.RedisPool.IdleCount() <= m.MaximumIdleConnection
}

func (m *Redis) DO(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := m.getConnection()
	defer conn.Close()
	reply, err = conn.Do(commandName, args)
	Assert(err)
	return
}

func (m *Redis) EXISTS(key string) bool {
	conn := m.getConnection()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if Assert(err) {
		return false
	}
	return exists
}

func (m *Redis) GET(key string) ([]byte, error) {
	conn := m.getConnection()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", key))
	if Assert(err) {
		return nil, err
	}
	return reply, nil
}

func (m *Redis) DELETE(key string) (bool, error) {
	conn := m.getConnection()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}

func (m *Redis) LIKE(key string) ([]string, error) {
	conn := m.getConnection()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if Assert(err) {
		return nil, err
	}
	return keys, nil
}

func (m *Redis) DELETEwithLIKE(key string) error {
	conn := m.getConnection()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if Assert(err) {
		return err
	}
	for _, key := range keys {
		_, err = m.DELETE(key)
		if Assert(err) {
			return err
		}
	}
	return nil
}

func (m *Redis) SETwithEXPIRE(key string, data interface{}, expire int) (err error) {
	conn := m.getConnection()
	defer conn.Close()
	value, err := json.Marshal(data)
	if Assert(err) {
		return
	}
	_, err = conn.Do("SET", key, value)
	if Assert(err) {
		return
	}
	_, err = conn.Do("EXPIRE", key, expire)
	if Assert(err) {
		return
	}
	return
}

func (m *Redis) FLUSH() error {
	conn := m.getConnection()
	defer conn.Close()
	err := conn.Flush()
	if Assert(err) {
		return err
	}
	return nil
}

func (m *Redis) SEND(commandName string, args ...interface{}) error {
	conn := m.getConnection()
	defer conn.Close()
	err := conn.Send(commandName, args...)
	if Assert(err) {
		return err
	}
	return nil
}

func (m *Redis) RECEIVE() (reply interface{}, err error) {
	conn := m.getConnection()
	defer conn.Close()
	result, err := conn.Receive()
	if Assert(err) {
		return result, err
	}
	return result, err
}
