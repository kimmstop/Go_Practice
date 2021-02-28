package Page

import (
	"io/ioutil"
)

type Pagedata struct {
	Title string
	Body []byte
}

func LoadPage(title string) (*Pagedata, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Pagedata{Title: title, Body: body}, nil
}