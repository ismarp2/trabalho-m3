package models

import "time"

type Usuario struct {
	ID             string    `json:"id"`
	Cpf            string    `json:"cpf"`
	Nome           string    `json:"nome"`
	DataNascimento time.Time `json:"data_nascimento"`
}
