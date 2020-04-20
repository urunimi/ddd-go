package repo

import (
	"github.com/urunimi/ddd-go/internal/pkg/notice"
)

type entityMapper struct{}

func (e entityMapper) dbNoticeToNotice(dbNotice *Notice) *notice.Notice {
	return &notice.Notice{
		ID:        dbNotice.ID,
		Title:     dbNotice.Title,
		Content:   dbNotice.Content,
		UserTypes: dbNotice.UserTypes,
		Lang:      dbNotice.Lang,
		CreatedAt: dbNotice.CreatedAt,
		UpdatedAt: dbNotice.UpdatedAt,
	}
}
