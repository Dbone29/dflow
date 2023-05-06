package rbac

import (
	"github.com/Dbone29/authority"
	"github.com/Dbone29/dflow/internal/database"
	"go.uber.org/zap"
)

type DflowRbac struct {
	logger *zap.Logger
	auth   *authority.Authority
}

func InitRbac(logger *zap.Logger, dflowDB *database.DflowDatabase) *DflowRbac {
	auth := authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           dflowDB.Database,
	})
	return &DflowRbac{
		logger: logger,
		auth:   auth,
	}
}
