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

func ObterLocacoes(c *gin.Context) {
	colecao := obterConnexaoLocacao()

	locacoes, err := colecao.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, locacoes)
}

func AdicionarLocacao(c *gin.Context) {
	var novaLocacao models.Locacao

	if err := c.BindJSON(&novaLocacao); err != nil {
		return
	}
	colecao := obterConnexaoLocacao()

	colecao.InsertOne(context.TODO(), novaLocacao)

	c.IndentedJSON(http.StatusOK, novaLocacao)
}

func AtualizarLocacao(c *gin.Context) {
	var locacao models.Locacao

	if err := c.BindJSON(&locacao); err != nil {
		return
	}
	colecao := obterConnexaoLocacao()

	colecao.UpdateByID(context.TODO(), locacao.ID, locacao)

	c.IndentedJSON(http.StatusOK, locacao)
}

func RemoverLocacao(c *gin.Context) {
	var locacao models.Locacao

	if err := c.BindJSON(&locacao); err != nil {
		return
	}
	colecao := obterConnexaoLocacao()

	colecao.DeleteOne(context.TODO(), locacao.ID)

	c.IndentedJSON(http.StatusOK, locacao)
}

func obterConnexaoLocacao() *mongo.Collection {
	client := connect()
	collection := client.Database("trabalho-m3").Collection("locacao")
	return collection
}
