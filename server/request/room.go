package request

type CreateRoom struct {
	Name   string `form:"name" validate:"required,excludesall= "`
	UserID string `form:"user_id" validate:"required"`
}