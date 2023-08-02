package performence

import "log"

func Gorunfunction() {
	protect(func() {
		a := 2
		b := 0
		c := a / b
		println(c)
	})
}

func protect(g func()) {
	defer func() {
		log.Println("recover")

		if x := recover(); x != nil {
			log.Println("run time panic: %v", x)
		}
	}()

	log.Println("start")
	g()
}
