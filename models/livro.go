package models

type Livro struct {
	ID                string `json:"id"`
	Titulo            string `json:"titulo"`
	Autor             string `json:"autor"`
	QuantidadePaginas int    `json:"quantidade_paginas"`
}
