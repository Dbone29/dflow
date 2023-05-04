package plugin

import (
	"github.com/Dbone29/dflow/pkg/events"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DflowPluginState struct {
	Logger       *zap.Logger
	EventManager *events.Event
	Database     *gorm.DB
}

type DflowPluginInfo struct {
	Name    string
	Version string
}

type DflowPlugin interface {
	Load(input DflowPluginState) error
	Unload(input DflowPluginState) error
	GetInfo() DflowPluginInfo
}
