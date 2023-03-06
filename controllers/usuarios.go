package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/isaquerchaves/DevBook/config"
	"github.com/isaquerchaves/DevBook/models"
)

func CriarUsuario(c *gin.Context) {
	var usuario struct {
		ID    uint
		Nome  string
		Nick  string
		Email string
		Senha string
	}

	c.Bind(&usuario)

	user := models.Usuario{Nome: usuario.Nome, Nick: usuario.Nick, Email: usuario.Email, Senha: usuario.Senha}

	// Validação (informações vazias)
	validacao := user.Preparar()
	if validacao != "" {
		c.JSON(400, gin.H{
			"": validacao,
		})
		return
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"": user,
	})
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")

	var user models.Usuario
	config.DB.First(&user, id)

	c.JSON(200, gin.H{
		"": user,
	})
}
func BuscarUsuarios(c *gin.Context) {
	var users []models.Usuario
	config.DB.Find(&users)

	c.JSON(200, gin.H{
		"": users,
	})
}
func AtualizarUsuarios(c *gin.Context) {
	id := c.Param("id")

	var userStruct struct {
		ID    uint
		Nome  string
		Nick  string
		Email string
		Senha string
	}

	c.Bind(&userStruct)

	var user models.Usuario
	config.DB.First(&user, id)

	config.DB.Model(&user).Updates(models.Usuario{
		Nome:  userStruct.Nome,
		Nick:  userStruct.Nick,
		Email: userStruct.Email,
		Senha: userStruct.Senha,
	})

	c.JSON(200, gin.H{
		"": user,
	})
}
func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Usuario{}, id)

	c.Status(200)
}
