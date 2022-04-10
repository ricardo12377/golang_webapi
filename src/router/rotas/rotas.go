package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

//representa todas as rotas
type Rota struct {
	URI                string
	Metodo             string
	Func               func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//coloca todas as rotas.
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Func).Methods(rota.Metodo)
	}

	return r
}
