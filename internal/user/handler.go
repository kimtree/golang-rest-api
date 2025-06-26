package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

// @Summary Create a user
// @Description This API will create a user
// @Tags        users
// @Accept json
// @Produce json
// @Param       request body user.CreateUserReq true "User Info"
// @Success 201 {object} user.User
// @Router /users [post]
func (h *Handler) CreateHandler(c *gin.Context) {
	var req CreateUserReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := User{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := h.repo.GetByEmail(u.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}

	if err := h.repo.Create(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetAllHandler(c *gin.Context) {
	var req GetAllUsersReq
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	page := max(1, req.Page)
	perPage := min(10, max(1, req.PageSize))
	offset := (page - 1) * perPage

	users, err := h.repo.GetAll(offset, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetByIDHandler(c *gin.Context) {
	var p UriParams
	if err := c.BindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.repo.GetByID(p.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateAllHandler(c *gin.Context) {
	var p UriParams
	if err := c.BindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req CreateUserReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := User{
		ID:    p.ID,
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := h.repo.Update(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdatePartialHandler(c *gin.Context) {
	var p UriParams
	if err := c.BindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req PatchUserReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ID 존재하지 않는 경우 이슈
	u, err := h.repo.GetByID(p.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Name != nil {
		u.Name = *req.Name
	}
	if req.Email != nil {
		u.Email = *req.Email
	}

	user, err := h.repo.Update(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteHandler(c *gin.Context) {
	var p UriParams
	if err := c.BindUri(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Delete(p.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
