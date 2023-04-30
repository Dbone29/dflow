package pipeline

import "github.com/gin-gonic/gin"

var InitApi = &Pipeline{}

type InitApiPayload struct {
	Gin *gin.Engine
}
