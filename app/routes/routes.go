package routes

import (
	v1 "github.com/anurag925/identity/app/handlers/api/v1"
	"github.com/anurag925/identity/core"
)

func Init() {
	server := core.Server()
	api := server.Group("api")
	apiV1 := api.Group("v1")

	applicationHandler := v1.NewApplicationHandler()
	apiV1.GET("/", applicationHandler.Index)
}
