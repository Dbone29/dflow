package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DflowApp struct {
	Logger    *zap.Logger
	Intl      *i18n.Bundle
	Database  *gorm.DB
	WebServer *gin.Engine
}
