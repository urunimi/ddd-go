package api

import (
	"github.com/jinzhu/gorm"
	"github.com/urunimi/ddd-go/internal/app/api/noticesvc"
	"github.com/urunimi/ddd-go/internal/pkg/notice"
	noticeRepo "github.com/urunimi/ddd-go/internal/pkg/notice/repo"
	"github.com/urunimi/gorest/core"
)

type apiApp struct {
	db *gorm.DB
}

func (ag *apiApp) Init() error {
	ag.db = getDatabase()
	return nil
}

func (ag *apiApp) RegisterRoute(driver *core.Engine) {
	nu := notice.NewUseCase(noticeRepo.New(ag.db))
	noticesvc.NewController(driver, nu)

	core.Logger().Infof("apiApp.RegisterRoute()")
}

func (ag *apiApp) Clean() error {
	return nil
}

var agApp *apiApp

// CreateAPIApp returns new core.App implementation.
func CreateAPIApp() core.App {
	agApp = &apiApp{}
	return core.App(agApp)
}
