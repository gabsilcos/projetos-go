package repo

//pacote que faz a conexão com o banco
//fazer o get no pacote de drivers do sql antes de rodar: go get -u github.com/go-sql-driver/mysql

import (
	"log"

	"github.com/jmoiron/sqlx"
	//github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
	_ "github.com/go-sql-driver/mysql" //o '_' antes do pacote é pra ele ser executado pela fç init, antes de tudo
)

//variavel singleton que armazena a conexao com o banco de dados
var db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (db *sqlx.DB, err error) {
	err = nil
	db, err = sqlx.Open("mysql", "root@tcp(localhost:3306)/cursodego?parseTime=true") //não esquecer da porta
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro na conexao: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro no ping na conexao: ", err.Error())
		return
	}
	return
}

//GetDBConnection Obtem a conexao com o banco de dados
func GetDBConnection() (localdb *sqlx.DB, err error) {
	if db == nil {
		db, err = AbreConexaoComBancoDeDadosSQL()
		if err != nil {
			log.Println("[GetDBConnection] Erro na conexao: ", err.Error())
			return
		}
	}
	err = db.Ping()
	if err != nil {
		log.Println("[GetDBConnection] Erro no ping na conexao: ", err.Error())
		return
	}
	localdb = db
	return
}
