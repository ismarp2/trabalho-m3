package models

import "time"

type Locacao struct {
	ID           string    `json:"id"`
	Livro        Livro     `json:"livro"`
	Usuario      Usuario   `json:"usuario"`
	DataLocacao  time.Time `json:"data_locacao"`
	PrazoEntrega int       `json:"prazo_entrega"`
}
