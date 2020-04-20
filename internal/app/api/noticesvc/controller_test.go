package noticesvc_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urunimi/ddd-go/internal/app/api/noticesvc"
	"github.com/urunimi/ddd-go/internal/pkg/notice"
	"github.com/urunimi/gorest/core"
	"github.com/urunimi/gorest/rest"
)

func (ts *ControllerTestSuite) TestController_GetNotice() {
	var mockNotices notice.Notices
	err := faker.FakeData(&mockNotices)
	ts.NoError(err)
	ts.usecase.On("GetBy", "PRO", "en", (*int64)(nil)).Return(mockNotices, nil)
	req := (&rest.Request{
		Method: http.MethodGet,
		URL:    "/notices",
		Params: &url.Values{
			"userType": []string{"PRO"},
			"locale":   []string{"en_US"},
		},
	}).Build()
	ctx, rec := ts.buildContextAndRecorder(req.GetHttpRequest())
	err = ts.controller.GetNotice(ctx)
	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.usecase.AssertExpectations(ts.T())
}

func (ts *ControllerTestSuite) TestController_PostNotice() {
	var mockNotice notice.Notice
	err := faker.FakeData(&mockNotice)
	mockNotice.ID = 0 // 새로운 Notice
	ts.NoError(err)
	var nPointer *notice.Notice
	ts.usecase.On("First", mock.Anything).Return(nPointer, errors.New("no item"))
	ts.usecase.On("Save", mock.Anything).Return(nil)
	req := (&rest.Request{
		Method: http.MethodPost,
		URL:    "/notices",
		Params: &url.Values{
			"title":    []string{mockNotice.Title},
			"content":  []string{mockNotice.Content},
			"language": []string{mockNotice.Lang},
		},
	}).Build()
	ctx, rec := ts.buildContextAndRecorder(req.GetHttpRequest())
	err = ts.controller.PostNotice(ctx)
	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)
	ts.usecase.AssertExpectations(ts.T())
}

func (ts *ControllerTestSuite) buildContextAndRecorder(httpRequest *http.Request) (ctx core.Context, rec *httptest.ResponseRecorder) {
	engine := core.NewEngine()
	rec = httptest.NewRecorder()
	ctx = engine.NewContext(httpRequest, rec)
	return
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(ControllerTestSuite))
}

type ControllerTestSuite struct {
	suite.Suite
	controller noticesvc.Controller
	usecase    *MockUseCase
}

func (ts *ControllerTestSuite) SetupTest() {
	engine := core.NewEngine()
	ts.usecase = new(MockUseCase)
	ts.controller = noticesvc.NewController(engine, ts.usecase)
}

var _ notice.UseCase = &MockUseCase{}

type MockUseCase struct {
	mock.Mock
}

func (u *MockUseCase) First(title string) (*notice.Notice, error) {
	ret := u.Called(title)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	} else {
		return ret.Get(0).(*notice.Notice), ret.Error(1)
	}
}

func (u *MockUseCase) GetBy(lang, userType string, lastId *int64) (notice.Notices, error) {
	ret := u.Called(userType, lang, lastId)

	var r0 notice.Notices
	if rf, ok := ret.Get(0).(func(string, string, *int64) notice.Notices); ok {
		r0 = rf(lang, userType, lastId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(notice.Notices)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *int64) error); ok {
		r1 = rf(lang, userType, lastId)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (u *MockUseCase) Save(n *notice.Notice) error {
	ret := u.Called(n)

	var r0 error
	if rf, ok := ret.Get(0).(func(*notice.Notice) error); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}
