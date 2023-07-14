package dataformatter

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func GorunJson() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}

	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)

	// file, _ := os.OpenFile("/resources/vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	// defer file.Close()
	// enc := json.NewEncoder(file)
	// err := enc.Encode(vc)
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Println("Error in encoding json")
	// }
	var vsObj = new(VCard)
	json.Unmarshal(js, vsObj)
	fmt.Println(vsObj.Addresses)
	fmt.Println(vsObj.FirstName)
	fmt.Println(vsObj.LastName)
	fmt.Println(vsObj.Remark)

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	var f interface{}
	err1 := json.Unmarshal(b, &f)
	fmt.Println(err1)

	m := f.(map[string]interface{})
	fmt.Println(m)

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)

		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}

	familyMember := FamilyMember{"Jack", 20, []string{"Tom", "Jerry"}}
	//famJson, err3 := json.Marshal(familyMember)
	// if err3 != nil {
	// 	fmt.Println(err3)
	// }

	var m1 FamilyMember
	json.Unmarshal(b, &m1)

	familyMember1 := FamilyMember{"Mack", 20, []string{"Lily", "Kites"}}
	fams := make([]FamilyMember, 0)
	fams = append(fams, familyMember, familyMember1)

	famsJson, _ := json.Marshal(fams)
	fmt.Print(string(famsJson))
}
