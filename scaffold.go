package scaffold

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
)

var ErrNotEnabled = errors.New("feature not enabled")

var _ GinApi = new(WebappScaffold)
var _ PostgresApi = new(WebappScaffold)

type WebappScaffold struct {
	g      *gin.Engine
	pgPool *pgxpool.Pool

	osFs   afero.Fs
	config WebappScaffoldConfig
}

func (w *WebappScaffold) GetPostgresPool() *pgxpool.Pool {
	if !w.config.PgConfig.Enable {
		panic(ErrNotEnabled)
	}
	return w.pgPool
}

func (w *WebappScaffold) GetGin() *gin.Engine {
	return w.g
}

func NewFromConfigFile(file string) (*WebappScaffold, error) {
	scaffold := new(WebappScaffold)
	scaffold.osFs = afero.NewOsFs()

	f, err := scaffold.osFs.Open(file)
	if err != nil {
		return nil, err
	}
	configData, err := afero.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(configData, &scaffold.config)
	if err != nil {
		return nil, err
	}

	if err := initGin(scaffold); err != nil {
		return nil, err
	}
	if err := initPg(scaffold); err != nil {
		return nil, err
	}

	return scaffold, nil
}

func initPg(scaffold *WebappScaffold) (err error) {
	if !scaffold.config.PgConfig.Enable {
		return nil
	}

	scaffold.pgPool, err = pgxpool.Connect(context.Background(), scaffold.config.PgConfig.PostgresConnectString)
	if err != nil {
		return err
	}

	return nil
}

func initGin(scaffold *WebappScaffold) error {
	scaffold.g = gin.New()
	return nil
}
