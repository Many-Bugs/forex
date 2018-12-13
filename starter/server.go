package starter

import (
	"errors"
	"fmt"
	"forex/api"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/pprof"
	ginSessions "github.com/gin-contrib/sessions"
	ginCookie "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	errInternetConnection = errors.New("Not connected to the network")
)

type Server struct {
	Engine
	Mode             string
	TLSCert          string
	IsNoCert         bool
	IsPerfamceCheck  bool
	Host             string
	Port             int
	Domain           string
	RequestTimeout   time.Duration
	TimeFormat       string
	TimeZone         string
	ServerExternalIP string
	CookieKey        string
	SessionsKey      string
}

type Engine struct {
	*gin.Engine
	HandlersFuncs []gin.HandlerFunc
}

// TODO: Map to Domain, later regester
func (m *Server) Builder(c *Content) error {

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

	m.newGinEngine(c)

	return nil
}

func (m *Server) ListenAndServe() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", m.Port),
		Handler:      m.Engine,
		ReadTimeout:  m.RequestTimeout,
		WriteTimeout: m.RequestTimeout,
	}
	go server.ListenAndServe()
}

func (m *Server) SecurityListenAndServe() {

}

func (m *Server) newGinEngine(c *Content) {
	m.Engine.Engine = gin.New()
	m.useMiddlewares()
	m.utilsServerSetting(c)
}

func (m *Server) utilsServerSetting(c *Content) {
	m.perfomanceCheck()
	gin.DefaultWriter = c.Logger.HTTPMessagesFile
	gin.DefaultErrorWriter = c.Logger.HTTPMessagesFile
	gin.SetMode(c.Server.Mode)
}

func (m *Server) perfomanceCheck() {
	if m.IsPerfamceCheck {
		pprof.Register(m.Engine.Engine, "debug/pprof")
	}
}

func (m *Server) useMiddlewares() {
	m.Engine.Use(
		api.CORS(),
		gin.Logger(),
		gin.Recovery(),
		ginSessions.Sessions(m.SessionsKey, ginCookie.NewStore([]byte(m.CookieKey))),
	)
	return
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

func (m *Server) Starter(c *Content) error {
	if m.IsNoCert {
		m.StartNoCert()
	}
	return nil
}

func (m *Server) Router(r Router) {
	m.Router(r)
}

func (m *Server) StartNoCert() {

	fmt.Fprintf(os.Stderr, "--- Started [:%d] ---\n", m.Port)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", m.Port),
		Handler:      m.Engine,
		ReadTimeout:  m.RequestTimeout,
		WriteTimeout: m.RequestTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error = %s\n", err.Error())
	}
}
