package noticesvc

type (
	GetNoticeRequest struct {
		Locale   string `form:"locale" validate:"required"`
		UserType string `form:"userType"`
		LastId   *int64 `form:"lastId"`
	}

	PostNoticeRequest struct {
		Title     string  `form:"title" validate:"required"`
		Content   string  `form:"content" validate:"required"`
		UserTypes *string `form:"userTypes"`
		Language  string  `form:"language" validate:"required"`
	}
)
