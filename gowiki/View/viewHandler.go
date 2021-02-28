package View

import (
	"net/http"
	"Page"
	"Template"
) 

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, _ := Page.LoadPage(title)
	Template.RenderTemplate(w, "view", p)
}
