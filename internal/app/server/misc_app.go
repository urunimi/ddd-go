package server

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/urunimi/gorest/core"
)

type miscApp struct {
}

func (ag *miscApp) Init() error {
	return nil
}

func (ag *miscApp) RegisterRoute(driver *core.Engine) {
	// driver.Renderer = &TemplateRegistry{
	// 	templates: template.Must(template.ParseGlob("./template/**/*.html")),
	// }
	driver.GET("/", ping)
	driver.GET("/ping", ping)
	driver.GET("/release", release)
	driver.GET("/panic", func(c core.Context) error {
		panic("this is test panic")
	})
	driver.POST("/error", func(c core.Context) error {
		return &echo.HTTPError{Code: http.StatusTeapot, Message: "this is test http error", Internal: errors.New("this is test error")}
	})
	driver.Static("/static", getProjectPath("/static"))
	driver.File("/favicon.ico", getProjectPath("static/img/favicon.ico"))
}

// TemplateRegistry Define the template registry struct
type TemplateRegistry struct {
	templates *template.Template
}

// Render Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c core.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (ag *miscApp) Clean() error {
	return nil
}

func release(c core.Context) error {
	err := c.Redirect(http.StatusMovedPermanently, "https://www.dropbox.com/sh/gyqrq3k0g0jnhdq/d20_XD4S0U")
	defer core.Logger().Infof("Releases() - UserAgent: %v", c.Request().UserAgent())
	return err
}

func ping(c core.Context) error {
	return c.JSON(http.StatusOK, "{code: 0, message: \"OK!\"}")
}

func getProjectPath(path string) string {
	return os.Getenv("GOPATH") + "/src/github.com/urunimi/ddd-go" + path
}
