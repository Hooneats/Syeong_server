package main

import (
	"github.com/Hooneats/Syeong_server/common/app"
	"github.com/Hooneats/Syeong_server/common/chiper"
	"github.com/Hooneats/Syeong_server/common/enum"
	"github.com/Hooneats/Syeong_server/common/flag"
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
	chiper.LoadCipherKey(config.ServerConfig.Mode)
	chiper.LoadCipherBlock()
	if err := config.JWTConfig.DecryptSalt(); err != nil {
		log.Fatal(err)
	}
	if err := config.DBConfig.DecryptURIAndDBName(); err != nil {
		log.Fatal(err)
	}

	//logger
	logger.LoadLogger(config.LogConfig)

	// model
	err := model.LoadMongoModel(config.DBConfig.URI)
	if err != nil {
		log.Fatal(err)
	}
	model.LoadMongoCollections(mongoCollectionNames, config.DBConfig.DBName)
	model.CreateIndexesInModels()
	model.InjectModelsMongoDependency(model.MongoCollection)

	// service
	service.InjectServicesDependency()

	// controller
	controller.InjectControllerDependency()

	// router
	router.SetAppRoute(
		router.NewGinRoute(config.ServerConfig.Mode),
	)

}

func main() {
	App.Run()
}
