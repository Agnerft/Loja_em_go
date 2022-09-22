package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Conexao struct {
	ConexaoLocal  string
	ConexaoHiroku string
}

var conn = Conexao{}

func Connection() *sql.DB {
	//conexaoBancoMake := "host=makesystem-production02.postgres.database.azure.com database=webcrm_vendas user=vendas-user password=Make@123$$ port=5432 sslmode=require"
	//conexaoHiroku := "host=ec2-34-199-68-114.compute-1.amazonaws.com database=d7kp419i1o9gr0 password=bed6d1caf8c5dcc3771c906ecc4ea6a1955a04852dd87c98b003e8578c8f604c user=vrguvrkqobonyj port=5432 sslmode=require"

	conexaoLocal := fmt.Sprintf("host=localhost database=postgres password=postgres user=postgres port=5432 sslmode=disable")
	db, err := sql.Open("postgres", conexaoLocal)
	if err != nil {
		panic(err.Error())
	}
	return db
}
