package repo

import (
	"github.com/user/banco_mongo/model"
	"gopkg.in/mgo.v2/bson"
)

//ObtemLocal retorna um local do MongoDB
func ObtemLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := SessaoMongo.Copy() //mongo não aceita pool de sessões, é preciso copiar
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("local") //sessao curso, coleção local
	err = colecao.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}

//SalvaLog registra a consulta ao Local
func SalvaLog(reg model.RegistroLog) (err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("logvisitas")
	err = colecao.Insert(reg) //como salvar coisas no mongo: '.Insert()' lol
	return
}
