package pluginmanager

import (
	"path/filepath"
	"plugin"

	dplugin "github.com/Dbone29/dflow/pkg/plugin"

	"go.uber.org/zap"
)

type PluginManager struct {
	logger    *zap.Logger
	pluginDir string
	plugins   []dplugin.DflowPlugin
}

func InitPluginManager(logger *zap.Logger, pluginDir string, plugins *[]dplugin.DflowPlugin) *PluginManager {
	var p []dplugin.DflowPlugin

	if plugins != nil {
		p = *plugins
	} else {
		p = make([]dplugin.DflowPlugin, 0)
	}

	return &PluginManager{
		logger:    logger,
		pluginDir: pluginDir,
		plugins:   p,
	}
}

func (pm *PluginManager) LoadPlugins() error {
	// Find all files in the plugin directory
	pluginFiles, err := filepath.Glob(filepath.Join(pm.pluginDir, "*.so"))
	if err != nil {
		pm.logger.Error("Failed to find plugins", zap.Error(err))
		return err
	}

	// Load and initialize each plugin
	for _, file := range pluginFiles {
		p, err := plugin.Open(file)
		if err != nil {
			pm.logger.Error("Failed to load plugin", zap.String("file", file), zap.Error(err))
			continue
		}

		symbol, err := p.Lookup("NewPlugin")
		if err != nil {
			pm.logger.Error("Failed to find Plugin symbol in plugin", zap.String("file", file), zap.Error(err))
			continue
		}

		newPluginFunc, ok := symbol.(func() dplugin.DflowPlugin)
		if !ok {
			pm.logger.Error("Invalid NewPlugin function in plugin", zap.String("file", file))
			continue
		}

		plugin := newPluginFunc()

		pm.plugins = append(pm.plugins, plugin)

	}

	return nil
}

func (pm *PluginManager) ActivatePlugins(state dplugin.DflowPluginState) error {
	for _, plugin := range pm.plugins {
		err := plugin.Load(state)

		if err != nil {
			pm.logger.Error("Failed to load plugin", zap.String("plugin", plugin.GetInfo().Name), zap.Error(err))
			continue
		}
	}

	return nil
}
