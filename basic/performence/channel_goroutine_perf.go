package performence

import (
	"fmt"
	"html/template"
)

func Gorun() {
	var strTempl = template.Must(template.New("TName").Parse("<h2></h2>"))
	fmt.Println(strTempl)
}
