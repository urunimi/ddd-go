package repo

import (
	"fmt"

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
	return r.mapper.toNoticeEntity(&dbNotice), nil
}

func (r *noticeRepo) Find(lang, userType string, lastID *int64) (notice.Notices, error) {

	query := fmt.Sprintf("(user_types LIKE '%%%s%%' or user_types = '') and lang = '%s'", userType, lang)

	if lastID == nil {
		query += fmt.Sprintf(" and id > %d", lastID)
	}

	dbNotices := make([]*Notice, 0)
	err := r.db.Where(query).Order("id desc").Find(&dbNotices).Error
	if err != nil {
		return nil, err
	}
	notices := make(notice.Notices, 0)
	for _, dbNotice := range dbNotices {
		notices = append(notices, r.mapper.toNoticeEntity(dbNotice))
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
