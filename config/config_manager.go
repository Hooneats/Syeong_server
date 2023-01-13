package config

import (
	"github.com/Hooneats/Syeong_server/common/flag"
	configDB "github.com/Hooneats/Syeong_server/config/dev/db"
	configJWT "github.com/Hooneats/Syeong_server/config/dev/jwt"
	configLog "github.com/Hooneats/Syeong_server/config/log"
	configServer "github.com/Hooneats/Syeong_server/config/server"
	"github.com/naoina/toml"
	"log"
	"os"
)

var DBConfig *configDB.DB
var JWTConfig *configJWT.JWT
var LogConfig *configLog.Log
var ServerConfig *configServer.Server

func LoadConfigs(pathMap map[string]*string) {

	dpath := pathMap[flag.DatabaseFlag.Name]
	DBConfig = NewConfig(*dpath, &configDB.DB{})

	jpath := pathMap[flag.JWTFlag.Name]
	JWTConfig = NewConfig(*jpath, &configJWT.JWT{})

	lpath := pathMap[flag.LogConfigFlag.Name]
	LogConfig = NewConfig(*lpath, &configLog.Log{})

	spath := pathMap[flag.ServerConfigFlag.Name]
	ServerConfig = NewConfig(*spath, &configServer.Server{})

}

func NewConfig[T any](path string, t *T) *T {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal("start app... does not exists config file in ", path)
	}

	if err := toml.NewDecoder(file).Decode(t); err != nil {
		log.Fatal("start app... toml decode, fail")
	}
	return t
}
