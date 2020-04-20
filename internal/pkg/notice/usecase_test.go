package notice_test

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/urunimi/ddd-go/internal/pkg/notice"
)

func (ts *UseCaseTestSuite) TestNoticeUseCase_First() {
	// Given
	var mockNotice notice.Notice
	err := faker.FakeData(&mockNotice)
	ts.NoError(err)
	ts.repo.On("First", mockNotice.Title).Return(&mockNotice, nil).Once()

	// When
	result, err := ts.usecase.First(mockNotice.Title)

	// Then
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
	ts.repo.On("Find", lang, userTypes, mock.Anything).Return(mockNotices, nil).Once()

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
	ts.repo.On("Save", &mockNotice).Return(nil).Once()

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
	repo    *mockRepo
	usecase notice.UseCase
}

func (ts *UseCaseTestSuite) SetupTest() {
	ts.repo = new(mockRepo)
	ts.usecase = notice.NewUseCase(ts.repo)
}

var _ notice.Repository = &mockRepo{}

type mockRepo struct {
	mock.Mock
}

func (r *mockRepo) First(title string) (*notice.Notice, error) {
	ret := r.Called(title)
	return ret.Get(0).(*notice.Notice), ret.Error(1)
}

func (r *mockRepo) Find(lang, userType string, lastID *int64) (notice.Notices, error) {
	ret := r.Called(lang, userType, lastID)
	return ret.Get(0).(notice.Notices), ret.Error(1)
}

func (r *mockRepo) Save(n *notice.Notice) error {
	return r.Called(n).Error(0)
}
