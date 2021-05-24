package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		mewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		mewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		mewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return 
	}

	c.Set(userCtx, userId)
}