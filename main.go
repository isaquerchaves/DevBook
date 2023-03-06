package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaquerchaves/DevBook/config"
	"github.com/isaquerchaves/DevBook/controllers"
	"github.com/isaquerchaves/DevBook/models"
)

func init() {
	config.ConnectToDb()
}

func main() {
	r := gin.Default()

	config.DB.AutoMigrate(&models.Usuario{})

	r.POST("/usuarios", controllers.CriarUsuario)
	r.GET("/usuarios", controllers.BuscarUsuarios)
	r.GET("/usuarios/:id", controllers.BuscarUsuario)
	r.PUT("/usuarios/:id", controllers.AtualizarUsuarios)
	r.DELETE("/usuarios/:id", controllers.DeletarUsuario)
	r.Run()
}
