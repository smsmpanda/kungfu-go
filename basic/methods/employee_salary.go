package methods

import (
	"fmt"
	"time"
)

type employee struct {
	salary float32
}

type myTime struct {
	time.Time //anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:4]
}

func (emp *employee) giveRaise(f float32) {
	emp.salary += emp.salary * f
	return
}

func Gorun2131() {
	emp := new(employee)
	emp.salary = 12.23112264

	emp.giveRaise(0.5)

	fmt.Printf("%f \n", emp.salary)

	m := myTime{time.Now()}
	// 调用匿名 Time 上的 String 方法
	fmt.Println("Full time now:", m.String())
	// 调用 myTime.first3Chars
	fmt.Println("First 4 chars:", m.first3Chars())
}
