package noticesvc

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/urunimi/ddd-go/internal/app/api/common/dto"
	"github.com/urunimi/ddd-go/internal/pkg/notice"
	"github.com/urunimi/gorest/core"
	"github.com/urunimi/gorest/rest"
)

// Controller type definition
type Controller struct {
	usecase notice.UseCase
}

// NewController returns new controller instance.
func NewController(e *core.Engine, au notice.UseCase) Controller {
	nc := Controller{au}
	e.GET("/notices", nc.GetNotice)
	e.POST("/notices", nc.PostNotice)

	return nc
}

var (
	langMap = map[string]bool{
		"en": true,
		"ko": true,
	}
)

func (nc *Controller) GetNotice(c core.Context) error {
	var req GetNoticeRequest
	if err := c.Bind(&req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	} else if err := c.Validate(&req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	lang := "en"
	if notices, err := nc.usecase.GetBy(lang, req.UserType, req.LastId); err != nil {
		return err
	} else {
		response := rest.Response{
			Code: dto.CodeOK,
			Result: map[string]interface{}{
				"notices":  notices,
				"interval": 7,
			},
		}

		core.Logger().Infof("GetNotice() - userType: %v, lang: %v, notices: %v", req.UserType, lang, len(notices))
		return c.JSON(http.StatusOK, response)
	}
}

func (nc *Controller) PostNotice(c core.Context) error {
	var req PostNoticeRequest
	if err := c.Bind(&req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	} else if err := c.Validate(&req); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
	}
	n, err := nc.usecase.First(req.Title)
	if err != nil {
		n = &notice.Notice{}
	}
	n.Content = req.Content
	n.Lang = req.Language
	n.UserTypes = req.UserTypes
	core.Logger().Infof("PostNotice() - req: %v, notice: %v", req, n)
	if err := nc.usecase.Save(n); err != nil {
		return err
	} else {
		response := rest.Response{
			Code: dto.CodeOK,
			Result: map[string]interface{}{
				"notice": n,
			},
		}

		return c.JSON(http.StatusOK, response)
	}
}
