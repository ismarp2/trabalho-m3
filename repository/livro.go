package repository

import (
	"context"
	"mongo-with-golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ObterLivros(c *gin.Context) {
	colecao := obterConnexaoLivro()

	livros, err := colecao.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, livros)
}

func AdicionarLivro(c *gin.Context) {
	var novoLivro models.Livro

	if err := c.BindJSON(&novoLivro); err != nil {
		return
	}
	colecao := obterConnexaoLivro()

	colecao.InsertOne(context.TODO(), novoLivro)

	c.IndentedJSON(http.StatusOK, novoLivro)
}

func AtualizarLivro(c *gin.Context) {
	var livro models.Livro

	if err := c.BindJSON(&livro); err != nil {
		return
	}
	colecao := obterConnexaoLivro()

	colecao.UpdateByID(context.TODO(), livro.ID, livro)

	c.IndentedJSON(http.StatusOK, livro)
}

func RemoverLivro(c *gin.Context) {
	var livro models.Livro

	if err := c.BindJSON(&livro); err != nil {
		return
	}
	colecao := obterConnexaoLivro()

	colecao.DeleteOne(context.TODO(), livro.ID)

	c.IndentedJSON(http.StatusOK, livro)
}

func obterConnexaoLivro() *mongo.Collection {
	client := connect()
	collection := client.Database("trabalho-m3").Collection("livro")
	return collection
}
