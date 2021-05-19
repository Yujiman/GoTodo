package handler

import (
	"net/http"

	"github.com/Yujiman/GoTodo/"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		mewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
		// todo 4:20
	}
}
func (h *Handler) signIn(c *gin.Context) {

}
