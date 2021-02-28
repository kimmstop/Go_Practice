package Save

import (
	"io/ioutil"
	"Page"
)


func Savepage(p *(Page.Pagedata)) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // 정상작동했다면 nil 반환
}