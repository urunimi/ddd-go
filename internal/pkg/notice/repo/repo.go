package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/urunimi/ddd-go/internal/pkg/notice"
)

type noticeRepo struct {
	db     *gorm.DB
	mapper entityMapper
}

var _ notice.Repository = &noticeRepo{}

func (r *noticeRepo) First(title string) (*notice.Notice, error) {
	dbNotice := Notice{Title: title}
	if err := r.db.Where(&dbNotice).First(&dbNotice).Error; err != nil {
		return nil, err
	}
	return r.mapper.dbNoticeToNotice(&dbNotice), nil
}

func (r *noticeRepo) Find(query string, args ...interface{}) (notice.Notices, error) {
	dbNotices := make([]*Notice, 0)
	err := r.db.Where(query, args...).Order("id desc").Find(&dbNotices).Error
	if err != nil {
		return nil, err
	}
	notices := make(notice.Notices, 0)
	for _, dbNotice := range dbNotices {
		notices = append(notices, r.mapper.dbNoticeToNotice(dbNotice))
	}
	return notices, nil
}

func (r *noticeRepo) Save(n *notice.Notice) error {
	return r.db.Save(&Notice{
		ID:        n.ID,
		Title:     n.Title,
		Content:   n.Content,
		UserTypes: n.UserTypes,
		Lang:      n.Lang,
	}).Error
}

func New(db *gorm.DB) notice.Repository {
	return &noticeRepo{db, entityMapper{}}
}
