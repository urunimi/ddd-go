package notice

import "fmt"

// UseCase interface definition
type UseCase interface {
	First(title string) (*Notice, error)
	GetBy(lang string, userType string, lastID *int64) (Notices, error)
	Save(notice *Notice) error
}

type usecase struct {
	repo Repository
}

func (u *usecase) First(title string) (*Notice, error) {
	return u.repo.First(title)
}

func (u *usecase) GetBy(lang, userType string, lastID *int64) (Notices, error) {
	var whereQuery string
	switch userType {
	case "DEV":
		whereQuery = ""
	default:
		whereQuery = fmt.Sprintf("(user_types LIKE '%%%s%%' or user_types = '') and lang = '%s' and ", userType, lang)
	}
	if lastID == nil {
		var id int64 = -1
		lastID = &id
	}
	return u.repo.Find(whereQuery+"id > ?", lastID)
}

func (u *usecase) Save(notice *Notice) (err error) {
	return u.repo.Save(notice)
}

// NewUseCase returns new UseCase implementation
func NewUseCase(noticeRepo Repository) UseCase {
	return &usecase{repo: noticeRepo}
}
