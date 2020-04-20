package notice

// Repository interface definition
type Repository interface {
	First(title string) (*Notice, error)
	Find(lang, userType string, lastID *int64) (Notices, error)
	Save(notice *Notice) error
}
