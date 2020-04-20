package notice

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
	return u.repo.Find(lang, userType, lastID)
}

func (u *usecase) Save(notice *Notice) (err error) {
	return u.repo.Save(notice)
}

// NewUseCase returns new UseCase implementation
func NewUseCase(noticeRepo Repository) UseCase {
	return &usecase{repo: noticeRepo}
}

var _ UseCase = &usecase{}
