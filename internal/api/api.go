package api

import (
	"fmt"
	"time"

	"github.com/Dbone29/dflow/pkg/pipeline"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DflowApi struct {
	logger *zap.Logger
	port   int
	gin    *gin.Engine
}

func InitApi(logger *zap.Logger, port int) *DflowApi {

	// Event inistalize api
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	//	r.Use(bundle)

	pl := pipeline.InitApi.Execute(pipeline.InitApiPayload{
		Gin: r,
	})

	r = pl.(pipeline.InitApiPayload).Gin

	// Event api inistalized

	return &DflowApi{
		logger: logger,
		port:   port,
		gin:    r,
	}
}

func (api *DflowApi) Start() error {
	api.logger.Info("Starting API server...", zap.Int("port", api.port))

	return api.gin.Run(":" + fmt.Sprint(api.port))
}
