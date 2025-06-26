package comment

type CreateCommentReq struct {
	UserID uint   `json:"userID" binding:"required"`
	Text   string `json:"text" binding:"required"`
}
