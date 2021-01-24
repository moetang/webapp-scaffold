package scaffold

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ScaffoldLifecycle interface {
	SyncStart() error
	Shutdown() error
}

type GinApi interface {
	GetGin() *gin.Engine
}

type PostgresApi interface {
	GetPostgresPool() *pgxpool.Pool
}
