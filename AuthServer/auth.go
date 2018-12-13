package AuthServer

import (
	"forex/library/document"
	"forex/starter"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthServer struct {
	*starter.Content
	apis []string
}

func (m *AuthServer) Starter() error {
	for _, model := range ModelAddrs() {
		m.Mysql.AutoMigrateByAddr(model)
	}
	m.TemplateRoutes()
	m.Server.Starter(m.Content)
	return nil
}

func (m *AuthServer) TemplateRoutes() (routes []gin.IRoutes) {

	group := m.Server.Engine.Group("/template")

	routes = []gin.IRoutes{
		m.TemplateHttpConnectionCheck(group),
		m.TemplateUsers(group),
	}
	return
}

/*
func (m *AuthServer) Template() gin.IRoutes {
	type response struct {
		Status string
	}
	doc := document.Doc{
		FilePath:   "",
		API:        "/template",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Response:   response{Status: "template"},
		Handler: func(context *gin.Context) {
			context.JSON(http.StatusOK, response{Status: "template"})
		},
	}
	return m.Server.Handle(doc.ToJSONResponse())
}
*/

func (m *AuthServer) TemplateHttpConnectionCheck(group *gin.RouterGroup) gin.IRoutes {
	type response struct {
		Status string
	}
	resp := response{Status: "Connection successful"}
	doc := document.Doc{
		FilePath:   "",
		API:        "/connectionCheck",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Response:   resp,
		Handler: func(context *gin.Context) {
			context.JSON(http.StatusOK, resp)
		},
	}
	return group.Handle(doc.ToJSONResponse())
}

func (m *AuthServer) TemplateUsers(group *gin.RouterGroup) gin.IRoutes {
	type response struct {
		Users []User
	}
	resp := response{
		Users: []User{
			User{
				Username: "template user 1",
				ParentID: 0,
			},
			User{
				Username: "template user 2",
				ParentID: 0,
			},
		},
	}
	doc := document.Doc{
		FilePath:   "",
		API:        "/user",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Response:   resp,
		Handler: func(context *gin.Context) {
			context.JSON(http.StatusOK, resp)
		},
	}
	return group.Handle(doc.ToJSONResponse())
}
