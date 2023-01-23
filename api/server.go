package api

import (
	"github.com/gin-gonic/gin"
	"github.com/martinconic/eth-events-indexer/config"
	"github.com/martinconic/eth-events-indexer/network"
	"github.com/martinconic/eth-events-indexer/storage"
	"github.com/martinconic/eth-events-indexer/utils/constants"
	"github.com/spf13/viper"
)

type Server struct {
	Router *gin.Engine

	NetworkClient *network.NetworkClient

	database storage.Database
}

var server *Server

func StartServer(v *viper.Viper) {
	server = &Server{}
	server.Initialize(v)
	server.Run(v.GetString(constants.ApiServer))
}

func (server *Server) Initialize(v *viper.Viper) {
	server.Router = gin.Default()
	server.NetworkClient = network.NewNetworkClient(config.GetNetworkConfigs(v))
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	err := server.Router.Run(":" + addr)
	if err != nil {
		panic("unable to start server!")
	}
}
