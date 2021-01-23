package scaffold

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
)

type WebappScaffold struct {
	g      *gin.Engine
	pgPool *pgxpool.Pool

	osFs   afero.Fs
	config WebappScaffoldConfig
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

func initPg(scaffold *WebappScaffold) error {
	return nil
}

func initGin(scaffold *WebappScaffold) error {
	return nil
}
