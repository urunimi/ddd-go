package notice

// Repository interface definition
type Repository interface {
	First(title string) (*Notice, error)
	Find(query string, args ...interface{}) (Notices, error)
	Save(notice *Notice) error
}
