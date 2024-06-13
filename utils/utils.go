package utils

import(
	"net/http"
	"html/template"
)

func HandleError(w http.ResponseWriter, err error)bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func HandleTemplate(w http.ResponseWriter, tmpl *template.Template){
	err := tmpl.Execute(w, nil)
	if err != nil {
		HandleError(w, err)
	}
}