package flag

import "flag"

type FlagCategory struct {
	Name    string
	Default string
	Usage   string
}

var (
	ServerConfigFlag = &FlagCategory{
		Name:    "server",
		Default: "./config/server/config.toml",
		Usage:   "toml file to use for server configuration",
	}

	LogConfigFlag = &FlagCategory{
		Name:    "log",
		Default: "./config/log/config.toml",
		Usage:   "toml file to use for log configuration",
	}

	DatabaseFlag = &FlagCategory{
		Name:    "db",
		Default: "./config/dev/db/config.toml",
		Usage:   "toml file to use for database configuration",
	}

	JWTFlag = &FlagCategory{
		Name:    "jwt",
		Default: "./config/dev/jwt/config.toml",
		Usage:   "toml file to use for jwt configuration",
	}
)

func (f *FlagCategory) Load() *string {
	return flag.String(f.Name, f.Default, f.Usage)
}
