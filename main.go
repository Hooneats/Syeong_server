package main

import (
	"github.com/Hooneats/Syeong_server/common/app"
	"github.com/Hooneats/Syeong_server/common/ciper"
	"github.com/Hooneats/Syeong_server/common/enum"
	"github.com/Hooneats/Syeong_server/common/flag"
	"github.com/Hooneats/Syeong_server/common/redis"
	"github.com/Hooneats/Syeong_server/config"
	"github.com/Hooneats/Syeong_server/controller"
	"github.com/Hooneats/Syeong_server/logger"
	"github.com/Hooneats/Syeong_server/model"
	"github.com/Hooneats/Syeong_server/router"
	"github.com/Hooneats/Syeong_server/service"
	"log"
)

var (
	App = app.NewApp()

	flags = []*flag.FlagCategory{
		flag.ServerConfigFlag,
		flag.LogConfigFlag,
		flag.DatabaseFlag,
		flag.JWTFlag,
		flag.RedisFlag,
	}

	mongoCollectionNames = []string{
		enum.UserCollectionName,
		enum.ContractCollectionName,
	}
)

func init() {
	flag.FlagsLoad(flags)
	config.LoadConfigs(flag.Flags)

	//Decrypt
	ciper.LoadCipherKey(config.ServerConfig.Mode)
	ciper.LoadCipherBlock()
	config.DecryptConfigs()

	//Redis
	if err := redis.LoadRedisClient(config.RedisConfig.DNS); err != nil {
		log.Fatal(err)
	}

	//logger
	logger.LoadLogger(config.LogConfig)

	// model
	if err := model.LoadMongoModel(config.DBConfig.URI, config.DBConfig.DBName, mongoCollectionNames); err != nil {
		log.Fatal(err)
	}
	model.CreateIndexesInModels()
	model.InjectModelsMongoDependency(model.MongoCollections)

	// service
	service.InjectServicesDependency()

	// controller
	controller.InjectControllerDependency()

	// router
	router.LoadRouter(config.ServerConfig.Mode)

}

func main() {
	App.Run()
}
