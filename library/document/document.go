package document

import (
	"github.com/gin-gonic/gin"
)

type Doc struct {
	ModuleName string
	FilePath   string
	API        string
	Method     string
	StatusCode int
	Handler    func(*gin.Context)
	Response   interface{}
}

func (m *Doc) ToJSONFile() {

}

func (m *Doc) ToJSONResponse() (method, api string, handler func(c *gin.Context)) {
	m.ToJSONFile()
	return m.Method, m.API, m.Handler
}
