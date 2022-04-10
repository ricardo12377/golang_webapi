package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSON retorna uma resposta em json para requisicao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {

		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}

}

//Erro retorna um erro em formato json
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `"json": "erro"`
	}{
		Erro: erro.Error(),
	})
}
