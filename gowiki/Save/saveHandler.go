package Save

import (
	"net/http"
	"Page"
) 




func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page.Pagedata{Title: title, Body: []byte(body)} // 저장할 페이지
	err := Savepage(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}