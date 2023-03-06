package models

import (
	"strings"

	"github.com/badoux/checkmail"
	"github.com/isaquerchaves/DevBook/seguranca.go"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Nome  string
	Nick  string
	Email string
	Senha string
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar() string {
	if erro := usuario.validar(); erro != "" {
		return erro
	}

	usuario.formatar()
	return ""
}

func (usuario *Usuario) validar() string {
	if usuario.Nome == "" {
		return "O Nome é obrigatório e não pode estar em branco"
	}
	if usuario.Nick == "" {
		return "O Nick é obrigatório e não pode estar em branco"
	}
	if usuario.Email == "" {
		return "O Email é obrigatório e não pode estar em branco"
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return "O e-mail inserido é invalido"
	}
	if usuario.Senha == "" {
		return "A Senha é obrigatório e não pode estar em branco"
	}

	return ""
}

// Remover espaços das extremidades
func (usuario *Usuario) formatar() string {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	senhaComHash, erro := seguranca.Hash(usuario.Senha)
	if erro != nil {
		return "Erro"
	}

	usuario.Senha = string(senhaComHash)

	return ""
}
