package utils

import "github.com/gin-gonic/gin"

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

func RespondError(c *gin.Context, status int, message string) {
	RespondJSON(c, status, gin.H{"error": message})
}
