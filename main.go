package main

import (
	"net/http"

	"github.com/list_vendas/routes"

	_ "github.com/lib/pq"
)

func main() {

	routes.Router()
	http.ListenAndServe(":8080", nil)
}
