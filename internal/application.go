package internal

import (
	"github.com/dendun-nf/gin-golang-web-test/internal/commons"
	"github.com/gin-gonic/gin"
)

type Application struct {
	engine    *gin.Engine
	endpoints []commons.IEndpoint
	// should have container of services
}

func NewApplication(engine *gin.Engine, endpoints []commons.IEndpoint) *Application {
	return &Application{
		engine:    engine,
		endpoints: endpoints,
	}
}

func (a *Application) MapEndpoints() {
	api := a.engine.Group("/api")
	if a.endpoints == nil {
		return
	}

	for _, endpoint := range a.endpoints {
		endpoint.MapEndpoint(api)
	}
}

func (a *Application) Run() {
	err := a.engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
