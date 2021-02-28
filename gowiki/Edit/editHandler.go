package Edit

import (
	"net/http"
	"Page"
	"Template"
) 

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := Page.LoadPage(title)
	if err != nil {
		p = &Page.Pagedata{Title: title}
	}
//	t, _ := template.ParseFiles("edit.html") //인자로 넘겨진 파일을 읽고 *template.Template 반환
//	t.Execute(w, p) // renderTemplate함수로 대체
	Template.RenderTemplate(w, "edit", p)
}