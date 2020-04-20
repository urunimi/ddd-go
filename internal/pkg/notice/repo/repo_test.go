package repo_test

import (
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"github.com/urunimi/ddd-go/internal/pkg/common/test"
	"github.com/urunimi/ddd-go/internal/pkg/notice"
	"github.com/urunimi/ddd-go/internal/pkg/notice/repo"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

//https://github.com/jirfag/go-queryset/blob/master/queryset/queryset_test.go
func (ts *RepoTestSuite) Test_First() {
	notices := getTestNotices(1)
	req := `SELECT * FROM "notices" WHERE ("notices"."title" = $1) ORDER BY "notices"."id" ASC LIMIT 1`
	ts.mock.ExpectQuery(ts.FixedFullRe(req)).
		WillReturnRows((&test.MockRowBuilder{}).Add(*notices[0]).Build())

	n, err := ts.repo.First(notices[0].Title)
	ts.NoError(err)
	ts.Equal(*(notices[0]), *n)
}

func (ts *RepoTestSuite) Test_Find() {
	notices := getTestNotices(2)
	req := "SELECT * FROM \"notices\" ORDER BY id desc"
	rowsBuilder := test.MockRowBuilder{}
	for i := range notices {
		rowsBuilder.Add(*notices[i])
	}
	ts.mock.ExpectQuery(ts.FixedFullRe(req)).
		WillReturnRows(rowsBuilder.Build())
	results, err := ts.repo.Find("")
	ts.NoError(err)
	for i := range results {
		ts.Equal(*notices[i], *results[i])
	}
}

func (ts *RepoTestSuite) Test_Save() {
	n := getTestNotices(1)[0]
	query := `UPDATE "notices" SET "title" = $1, "content" = $2, "user_types" = $3, "lang" = $4, "updated_at" = $5 WHERE "notices"."id" = $6`
	args := []driver.Value{n.Title, sqlmock.AnyArg(), sqlmock.AnyArg(), n.Lang, sqlmock.AnyArg(), sqlmock.AnyArg()}
	ts.mock.ExpectBegin().WillReturnError(nil)
	ts.mock.ExpectExec(ts.FixedFullRe(query)).
		WithArgs(args...).
		WillReturnResult(sqlmock.NewResult(1, 1))
	ts.mock.ExpectCommit().WillReturnError(nil)
	err := ts.repo.Save(n)
	ts.NoError(err)
}

func getTestNotices(num int) []*notice.Notice {
	notices := make([]*notice.Notice, 0)
	for i := 1; i < num+1; i++ {
		n := &notice.Notice{
			ID:    int64(i),
			Title: fmt.Sprintf("Title - %d", i),
			Lang:  "en",
		}
		notices = append(notices, n)
	}
	return notices
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}

type RepoTestSuite struct {
	*test.GormHelper
	suite.Suite
	mock sqlmock.Sqlmock
	db   *gorm.DB
	repo notice.Repository
}

func (ts *RepoTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	ts.NoError(err)
	ts.mock = mock
	ts.db, err = gorm.Open("postgres", db)
	ts.NoError(err)
	ts.db.LogMode(true)
	ts.db = ts.db.Set("gorm:update_column", true)
	ts.repo = repo.New(ts.db)
}

func (ts *RepoTestSuite) AfterTest(suiteName, testName string) {
	ts.db.Close()
}
