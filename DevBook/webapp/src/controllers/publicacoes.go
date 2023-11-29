package controllers

import (
	"fmt"
	"net/http"
)

// CriarPublicacao chama a API para cadastar a publicaçao no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Criando Publicação")
}

/*	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErrorAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIURL)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErrorAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {

		respostas.TratarStatusCodeDeErro(w, response)
		return

	}

	respostas.JSON(w, response.StatusCode, nil)

}
*/
