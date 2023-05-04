/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Dbone29/dflow/internal/api"
	"github.com/Dbone29/dflow/internal/config"
	"github.com/Dbone29/dflow/internal/database"
	"github.com/Dbone29/dflow/internal/intl"
	"github.com/Dbone29/dflow/internal/log"
	pluginmanager "github.com/Dbone29/dflow/internal/plugin-manager"
	"github.com/Dbone29/dflow/internal/storage"
	"github.com/Dbone29/dflow/pkg/plugin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
		pm := pluginmanager.InitPluginManager(logger, "plugins")

		err = pm.LoadPlugins()
		if err != nil {
			logger.Error("Failed to load plugins", zap.Error(err))
		}

		err = pm.ActivatePlugins(plugin.DflowPluginState{
			Logger:   logger,
			Database: db,
		})
		if err != nil {
			logger.Error("Failed to activate plugins", zap.Error(err))
		}

		// init api server
		apiServer := api.InitApi(logger, 8080)

		// start api server
		logger.Info("starting server...")
		apiServer.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
