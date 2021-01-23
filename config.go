package scaffold

type WebappScaffoldConfig struct {
	GinConfig struct {
	} `toml:"gin"`
	PgConfig struct {
		Enable                bool   `toml:"enable" default:"true"`
		PostgresConnectString string `toml:"connstring"`
	} `toml:"datasource.postgres"`
}
