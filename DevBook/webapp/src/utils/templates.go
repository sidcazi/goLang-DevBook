package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// Carregar Templates insere os templates HTML na variavel template
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// Executar template renderiza uma pagina htm na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
