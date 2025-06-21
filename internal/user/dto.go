package user

type UriParams struct {
	ID uint `uri:"id" binding:"required"`
}

type CreateUserReq struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type PatchUserReq struct {
	Name  *string `json:"name" binding:"omitempty"`
	Email *string `json:"email" binding:"omitempty,email"`
}
