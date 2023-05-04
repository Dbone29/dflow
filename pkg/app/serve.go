package app

import (
	"github.com/Dbone29/dflow/internal/api"
	"github.com/Dbone29/dflow/internal/config"
	"github.com/Dbone29/dflow/internal/database"
	"github.com/Dbone29/dflow/internal/intl"
	"github.com/Dbone29/dflow/internal/log"
	pluginmanager "github.com/Dbone29/dflow/internal/plugin-manager"
	"github.com/Dbone29/dflow/internal/storage"
	dplugin "github.com/Dbone29/dflow/pkg/plugin"
	"go.uber.org/zap"
)

func Serve(plugins *[]dplugin.DflowPlugin) {
	// init logger
	logger, err := log.InitLogger()
	if err != nil {
		panic(err)
	}

	// load configs
	logger.Info("loading configs...")
	cf := config.LoadConfig()

	intl.InitIntl( /*cf.Main.Language*/ )

	// init database
	logger.Info("init database...")
	db := database.InitDatabase(logger, &cf.Main.Database)

	// init storage
	logger.Info("init storage...")
	storage.InitStorage(logger, &cf.Main.Storage)

	// init plugin manager
	logger.Info("loading plugins...")
	pm := pluginmanager.InitPluginManager(logger, "plugins", plugins)

	err = pm.LoadPlugins()
	if err != nil {
		logger.Error("Failed to load plugins", zap.Error(err))
	}

	err = pm.ActivatePlugins(dplugin.DflowPluginState{
		Logger:   logger,
		Database: db.Database,
	})
	if err != nil {
		logger.Error("Failed to activate plugins", zap.Error(err))
	}

	// init api server
	apiServer := api.InitApi(logger, 8080)

	// start api server
	logger.Info("starting server...")
	apiServer.Start()
}