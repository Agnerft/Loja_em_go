package controlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/list_vendas/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "index", produtos)

	/*prdutos := []Produto{
		{Nome: "Óculos", Descricao: "De grau", Preco: 89.90, Quantidade: 2},
		{"Ténis", "Tamanho 38", 109.90, 4},
	}*/
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão de erro:", err)
		}
		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConv, quantidadeConv)
	}

	http.Redirect(w, r, "/inicio", http.StatusMovedPermanently)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "/inicio", http.StatusMovedPermanently)

}

func Edit(w http.ResponseWriter, r *http.Request) {

	idProduto := r.URL.Query().Get("id")

	produto := models.SelecProdutoById(idProduto)

	temp.ExecuteTemplate(w, "edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idAjustado, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro no conversão do Id:", err.Error())
		}

		precoAjustado, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeAjustada, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.EditProduto(idAjustado, quantidadeAjustada, descricao, nome, precoAjustado)
	}

	http.Redirect(w, r, "/inicio", http.StatusMovedPermanently)

}

/*func Banco(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "escBanco", nil)
}*/
