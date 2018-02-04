package model

import "database/sql"

//esse modelo de struct espelha os campos do elemnto no sql: os dados passam por ela antes de serem levados
//ao banco

//Local armazena os dados da localidade pelo seu codigo telefonico
type Local struct {
	Pais             string         `json:"pais" db:"country"`
	Cidade           sql.NullString `json:"cidade" db:"city"`
	CodigoTelefonico int            `json:"cod_telefone" db:"telcode"`
}
