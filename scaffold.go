package scaffold

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/moetang/webapp-scaffold/frmgin"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pelletier/go-toml"
	"github.com/spf13/afero"
)

var ErrNotEnabled = errors.New("feature not enabled")

var _ ScaffoldLifecycle = new(WebappScaffold)
var _ GinApi = new(WebappScaffold)
var _ PostgresApi = new(WebappScaffold)

type WebappScaffold struct {
	g      *gin.Engine
	pgPool *pgxpool.Pool

	osFs   afero.Fs
	config WebappScaffoldConfig
}

func (w *WebappScaffold) SyncStart() error {
	if err := startPg(w); err != nil {
		return err
	}
	if err := startGin(w); err != nil {
		return err
	}

	return w.g.Run(w.config.GinConfig.Listen)
}

func startPg(scaffold *WebappScaffold) (err error) {
	if !scaffold.config.PgConfig.Enable {
		return
	}

	scaffold.pgPool, err = pgxpool.Connect(context.Background(), scaffold.config.PgConfig.PostgresConnectString)
	if err != nil {
		return
	}
	return
}

func (w *WebappScaffold) Shutdown() error {
	w.pgPool.Close()
	//FIXME need to implement
	panic("implement me")
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

	if scaffold.config.GlobalConfig.Debug {
		log.Println("[DEBUG] config loaded. Content:")
		log.Println("[DEBUG] ", scaffold.config)
	}

	if err := initGin(scaffold); err != nil {
		return nil, err
	}
	if err := initPg(scaffold); err != nil {
		return nil, err
	}

	return scaffold, nil
}

func ReadCustomConfig(file string, s interface{}) error {
	fs := afero.NewOsFs()

	f, err := fs.Open(file)
	if err != nil {
		return err
	}
	configData, err := afero.ReadAll(f)
	if err != nil {
		return err
	}

	return toml.Unmarshal(configData, s)
}

func initPg(scaffold *WebappScaffold) (err error) {
	return nil
}

func initGin(scaffold *WebappScaffold) error {
	if scaffold.config.GinConfig.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	scaffold.g = gin.New()
	// init builtin func
	frmgin.InitBuiltinFunc(scaffold.g)

	return nil
}

func startGin(scaffold *WebappScaffold) (err error) {
	for _, v := range scaffold.config.GinConfig.HtmlGlobPaths {
		scaffold.g.LoadHTMLGlob(v)
	}

	// debug mode: reload html glob files
	if !scaffold.config.GinConfig.ReleaseMode {
		go func() {
			ti := time.NewTicker(5 * time.Second)
			for {
				select {
				case <-ti.C:
					f := func() {
						defer func() {
							e := recover()
							if e != nil {
								log.Println("[ERROR] reload html glob fail:", e)
							}
						}()
						for _, v := range scaffold.config.GinConfig.HtmlGlobPaths {
							scaffold.g.LoadHTMLGlob(v)
						}
					}
					f()
					log.Println("[DEBUG] html glob refreshed.")
				}
			}
		}()
	}

	return
}
