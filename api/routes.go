package api

func (s *Server) InitializeRoutes() {
	//"0x3845badAde8e6dFF049820680d1F14bD3903a5d0"
	s.Router.GET("/api/contracts", GetSmartContracts)
	s.Router.GET("/api/events/:address", GetSmartContractEvents)
	s.Router.GET("/api/indexed/:id", GetIndexedEvents)
	s.Router.POST("/api/events/add/:address", AddSmartContract)
}
