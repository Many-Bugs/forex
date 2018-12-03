package starter

import (
	"errors"
	"net"
	"strconv"
	"time"

	ginSessions "github.com/gin-contrib/sessions"
	ginCookie "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	errInternetConnection = errors.New("Not connected to the network")
)

type Server struct {
	Mode             string
	Host             string
	Port             int
	Domain           string
	RequestTimeout   time.Duration
	TimeFormat       string
	TimeZone         string
	ServerExternalIP string
	CookieKey        string
	Engine           *gin.Engine
}

// TODO: Map to Domain, later regester
func (m *Server) Builder() error {

	ip, err := getLocalExternalIP()
	if err != nil {
		return err
	} else {
		m.ServerExternalIP = ip
	}

	for {
		if !checkPortAvailable(strconv.Itoa(m.Port)) {
			m.Port++
		} else {
			break
		}
	}

	m.RequestTimeout = m.RequestTimeout * time.Second
	if local, err := time.LoadLocation(m.TimeZone); err != nil {
		time.Local = time.UTC
	} else {
		time.Local = local
	}

	m.Engine = gin.New()
	store := ginCookie.NewStore([]byte(m.CookieKey))
	m.Engine.Use(ginSessions.Sessions("many_bugs_forex_session", store))

	return nil
}

func checkPortAvailable(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

func getLocalExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errInternetConnection
}
