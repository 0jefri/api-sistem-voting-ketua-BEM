package delivery

import (
	"github.com/api-voting/internal/app/delivery/routes"
	"github.com/gin-gonic/gin"
)

type application struct {
	engine *gin.Engine
}

func (app *application) Run() {
	if err := routes.SetupRouter(app.engine); err != nil {
		panic("Application error")
	}
}

func Server() *application {
	router := gin.Default()

	return &application{
		engine: router,
	}

}
