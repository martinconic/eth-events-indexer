package api

func (f *Server) InitializeRoutes() {
	f.Router.GET("/api/events", GetSmartContractEvents)
}
