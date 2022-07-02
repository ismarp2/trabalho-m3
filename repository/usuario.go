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

func ObterUsuarios(c *gin.Context) {
	colecao := obterConnexaoUsuario()

	usuarios, err := colecao.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, usuarios)
}

func AdicionarUsuario(c *gin.Context) {
	var novoUsuario models.Usuario

	if err := c.BindJSON(&novoUsuario); err != nil {
		return
	}
	colecao := obterConnexaoUsuario()

	colecao.InsertOne(context.TODO(), novoUsuario)

	c.IndentedJSON(http.StatusOK, novoUsuario)
}

func AtualizarUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.BindJSON(&usuario); err != nil {
		return
	}
	colecao := obterConnexaoUsuario()

	colecao.UpdateByID(context.TODO(), usuario.ID, usuario)

	c.IndentedJSON(http.StatusOK, usuario)
}

func RemoverUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.BindJSON(&usuario); err != nil {
		return
	}
	colecao := obterConnexaoUsuario()

	colecao.DeleteOne(context.TODO(), usuario.ID)

	c.IndentedJSON(http.StatusOK, usuario)
}

func obterConnexaoUsuario() *mongo.Collection {
	client := connect()
	collection := client.Database("trabalho-m3").Collection("usuario")
	return collection
}
