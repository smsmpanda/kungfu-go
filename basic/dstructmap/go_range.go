package dstructmap

import (
	"fmt"
)

func Gorun1321() {
	var GDPMap = map[string]float64{"beijing": 1415151.12, "shanghai": 14123123.12}
	for key, value := range GDPMap {
		fmt.Printf("%s's gdp is %f.\n", key, value)
	}

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	var map_days = make(map[string]int, 0)
	map_days["Mon"] = 1
	//map_days["Tues"] = 2
	map_days["Wed"] = 3
	map_days["Thurs"] = 4
	map_days["Frai"] = 5
	map_days["Stat"] = 6
	map_days["Sun"] = 7
	if _, ok := map_days["Tues"]; ok {
		fmt.Println("Tues is ating.")
	} else {
		fmt.Println("Tues isn't ating.")
	}
}
