package models

import (
	"fmt"

	"github.com/list_vendas/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.Connection()

	abrindoConexao, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for abrindoConexao.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = abrindoConexao.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		//fmt.Println(p)
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer db.Close()

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.Connection()

	criandoNovo, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	criandoNovo.Exec(nome, descricao, preco, quantidade)

	defer db.Close()

}

func DeletaProduto(id string) {
	db := db.Connection()

	deletandoProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletandoProduto.Exec(id)

	defer db.Close()
}

func SelecProdutoById(id string) Produto {
	db := db.Connection()

	editandoProduto, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	ProdutoAtt := Produto{}

	for editandoProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = editandoProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		ProdutoAtt.Id = id
		ProdutoAtt.Nome = nome
		ProdutoAtt.Descricao = descricao
		ProdutoAtt.Preco = preco
		ProdutoAtt.Quantidade = quantidade

		fmt.Print(ProdutoAtt)
	}

	defer db.Close()

	return ProdutoAtt

}

func EditProduto(id, quantidade int, descricao, nome string, preco float64) {
	db := db.Connection()

	updateProduto, err := db.Prepare("update produtos set nome=$2, descricao=$3, preco=$4, quantidade=$5 where id=$1")
	if err != nil {
		panic(err.Error())
	}

	updateProduto.Exec(id, nome, descricao, preco, quantidade)

}
