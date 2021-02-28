package Template

import (
	"html/template" // hard coded된 html대신 html 파일을 따로 보관하게 해주는 라이브러리 사용
	"net/http"
	"Page"
)

var templates = template.Must(template.ParseFiles("./html/edit.html", "./html/view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *(Page.Pagedata)) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
