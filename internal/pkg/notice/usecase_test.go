package notice_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urunimi/ddd-go/internal/pkg/notice"
)

func (ts *UseCaseTestSuite) TestNoticeUseCase_First() {
	var mockNotice notice.Notice
	err := faker.FakeData(&mockNotice)
	ts.NoError(err)
	ts.repo.On("First", mockNotice.Title).Return(&mockNotice, nil)
	result, err := ts.usecase.First(mockNotice.Title)
	ts.NoError(err)
	ts.Equal(result.Title, mockNotice.Title)
	ts.Equal(result.Content, mockNotice.Content)
	ts.repo.AssertExpectations(ts.T())
}

func (ts *UseCaseTestSuite) TestNoticeUseCase_GetBy() {
	var mockNotices notice.Notices
	lang, userTypes := "en", "PRO"
	err := faker.FakeData(&mockNotices)
	ts.NoError(err)
	for _, mn := range mockNotices {
		mn.Lang = lang
		mn.UserTypes = &userTypes
	}
	ts.repo.On("Find", `(user_types LIKE '%PRO%' or user_types = '') and lang = 'en' and id > ?`, mock.Anything).Return(mockNotices, nil)
	results, err := ts.usecase.GetBy("en", "PRO", nil)
	ts.NoError(err)
	ts.Equal(mockNotices, results)
	ts.repo.AssertExpectations(ts.T())
}

func (ts *UseCaseTestSuite) TestNoticeUseCase_Save() {
	var mockNotice notice.Notice
	err := faker.FakeData(&mockNotice)
	ts.NoError(err)
	mockNotice.ID = 0
	ts.repo.On("Save", &mockNotice).Return(nil)
	err = ts.usecase.Save(&mockNotice)
	ts.NoError(err)
	ts.repo.AssertExpectations(ts.T())
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(UseCaseTestSuite))
}

var (
	_ suite.SetupTestSuite = &UseCaseTestSuite{}
)

type UseCaseTestSuite struct {
	suite.Suite
	repo    *MockRepository
	usecase notice.UseCase
}

func (ts *UseCaseTestSuite) SetupTest() {
	ts.repo = new(MockRepository)
	ts.usecase = notice.NewUseCase(ts.repo)
}

var _ notice.Repository = &MockRepository{}

type MockRepository struct {
	mock.Mock
}

func (r *MockRepository) First(title string) (*notice.Notice, error) {
	ret := r.Called(title)
	return ret.Get(0).(*notice.Notice), ret.Error(1)
}

func (r *MockRepository) Find(query string, args ...interface{}) (notice.Notices, error) {
	ret := r.Called(query, args)
	return ret.Get(0).(notice.Notices), ret.Error(1)
}

func (r *MockRepository) Save(n *notice.Notice) error {
	return r.Called(n).Error(0)
}
