package plugin

import (
	"github.com/Dbone29/dflow/internal/database"
	"github.com/Dbone29/dflow/pkg/events"
	"go.uber.org/zap"
)

type DflowPluginState struct {
	Logger       *zap.Logger
	EventManager *events.Event
	Database     *database.DflowDatabase
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
