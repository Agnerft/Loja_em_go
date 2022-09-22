package routes

import (
	"net/http"

	"github.com/list_vendas/controlers"
)

func Router() {
	//http.HandleFunc("/", controlers.Banco)
	http.HandleFunc("/inicio", controlers.Index)
	http.HandleFunc("/new", controlers.New)
	http.HandleFunc("/insert", controlers.Inserir)
	http.HandleFunc("/delete", controlers.Delete)
	http.HandleFunc("/edit", controlers.Edit)
	http.HandleFunc("/update", controlers.Update)
}
