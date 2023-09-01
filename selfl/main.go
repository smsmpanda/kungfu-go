package selfl

import (
	"fmt"
	"myexample/basic/trans"
	"myexample/basic/trans/model"
	selflmodel "myexample/selfl/model"
)

func Gorun() {
	fmt.Printf("trans.Pi: %v\n", trans.Pi)

	f := &selflmodel.PhoneFactory{}
	p := f.GreatePhone("HUAWEI", "white")

	fmt.Printf("p.GetColor(): %v\n", p.GetColor())
	fmt.Printf("p: %v\n", p)

	fmt.Printf("model.Big: %v\n", model.Big)
}
