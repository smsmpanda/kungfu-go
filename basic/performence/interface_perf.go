package performence

import "fmt"

type ISelf interface {
	getSelf() string
	getSelies(itmes ...interface{})
}

type Me struct {
}

func (me *Me) getSelf(name string) string {
	return name
}

func (me *Me) getSelies(itmes ...interface{}) {
	for i, x := range itmes {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}

func GorunInterfaces() {
	me := &Me{}
	result := me.getSelf("this is interface")
	fmt.Println(result)

	me.getSelies(1, 2, 3, 5)
}
