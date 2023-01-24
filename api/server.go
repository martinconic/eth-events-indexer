package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/martinconic/eth-events-indexer/config"
	"github.com/martinconic/eth-events-indexer/network"
	"github.com/martinconic/eth-events-indexer/storage"
	"github.com/martinconic/eth-events-indexer/storage/psqldb"
	"github.com/martinconic/eth-events-indexer/utils/constants"
	"github.com/spf13/viper"
)

type Server struct {
	Router *gin.Engine

	NetworkClient *network.NetworkClient

	db storage.Database
}

var server *Server

func StartServer(v *viper.Viper) {
	server = &Server{}
	server.Initialize(v)
	server.Run(v.GetString(constants.ApiServer))
}

func (server *Server) Initialize(v *viper.Viper) {
	var err error
	server.Router = gin.Default()
	server.NetworkClient, err = network.NewNetworkClient(config.GetNetworkConfigs(v))
	if err != nil {
		log.Println(err)
	}
	server.db, err = psqldb.NewDatabase(config.GetPostgresConfig(v))
	if err != nil {
		log.Println(err)
	}
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	err := server.Router.Run(":" + addr)
	if err != nil {
		panic("unable to start server!")
	}
}
