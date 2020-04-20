package notice

import (
	"encoding/json"
	"time"
)

// RFC3339Milli is time format
const RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"

// Notice is
type Notice struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserTypes *string   `json:"-"`
	Lang      string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// MarshalJSON overrides
func (notice *Notice) MarshalJSON() ([]byte, error) {
	type Alias Notice
	return json.Marshal(&struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		*Alias
	}{
		CreatedAt: notice.CreatedAt.Format(RFC3339Milli),
		UpdatedAt: notice.UpdatedAt.Format(RFC3339Milli),
		Alias:     (*Alias)(notice),
	})
}

// Notices is
type Notices []*Notice

func (slice Notices) Len() int {
	return len(slice)
}

func (slice Notices) Less(i, j int) bool {
	return slice[i].UpdatedAt.After(slice[j].UpdatedAt)
}

func (slice Notices) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
