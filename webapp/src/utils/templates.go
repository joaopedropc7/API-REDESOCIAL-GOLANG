package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Carregar templates insere os templates HTML na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate renderiza uma pagina HTML na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
