package api

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func (server *Server) Initialize() {
	server.Router = gin.Default()

	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	err := server.Router.Run(":" + addr)
	if err != nil {
		panic("unable to start server!")
	}
}
