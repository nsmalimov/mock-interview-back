package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorizationRequest struct {
	Email string `json:"email"`
}

func Authorization(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("error when try io.ReadAll, err: %s", err))
		return
	}

	authR := &AuthorizationRequest{}

	if err := json.Unmarshal(jsonData, &authR); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("error when try json.Unmarshal, err: %s", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": authR.Email,
	})
}
