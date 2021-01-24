package scaffold

type WebappScaffoldConfig struct {
	GlobalConfig struct {
		Debug bool `toml:"debug"`
	} `toml:"global"`
	GinConfig struct {
		ReleaseMode   bool     `toml:"release_mod"`
		HtmlGlobPaths []string `toml:"html_glob_paths"`
		Listen        string   `toml:"listen" default:":6001"`
	} `toml:"gin"`
	PgConfig struct {
		Enable                bool   `toml:"enable" default:"false"`
		PostgresConnectString string `toml:"connstring"`
	} `toml:"postgres"`
}
