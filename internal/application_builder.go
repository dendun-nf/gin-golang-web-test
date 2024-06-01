package internal

import (
	"github.com/dendun-nf/gin-golang-web-test/internal/commons"
	"github.com/gin-gonic/gin"
)

type ApplicationBuilder struct {
	engine    *gin.Engine
	endpoints []commons.IEndpoint
}

func NewApplicationBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{
		engine: gin.Default(),
	}
}

func (a *ApplicationBuilder) AddEndpoints(endpoints ...commons.IEndpoint) *ApplicationBuilder {
	a.endpoints = append(a.endpoints, endpoints...)
	return a
}

func (a *ApplicationBuilder) Build() *Application {
	return NewApplication(a.engine, a.endpoints)
}
