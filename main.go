package main

import (
	"mongo-with-golang/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/livro", repository.ObterLivros)
	router.POST("/livro", repository.AdicionarLivro)
	router.PUT("/livro", repository.AtualizarLivro)
	router.DELETE("/livro", repository.RemoverLivro)

	router.GET("/locacao", repository.ObterLocacoes)
	router.POST("/locacao", repository.AdicionarLocacao)
	router.PUT("/locacao", repository.AtualizarLocacao)
	router.DELETE("/locacao", repository.RemoverLocacao)

	router.GET("/usuario", repository.ObterUsuarios)
	router.POST("/usuario", repository.AdicionarUsuario)
	router.PUT("/usuario", repository.AtualizarUsuario)
	router.DELETE("/usuario", repository.RemoverUsuario)

	router.Run("localhost:8080")
}
