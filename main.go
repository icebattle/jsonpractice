package main

import (
	"encoding/json"
	"fmt"

	"github.com/icebattle/jsonpractice/formats"
)

func main() {
	p := formats.Person{"Morning", "Evening"}
	fmt.Println(p)
	b, _ := json.Marshal(p)
	fmt.Println(b)

	var mp formats.Person
	json.Unmarshal(b, &mp)
	fmt.Println(mp)

	var uf interface{}
	json.Unmarshal(b, &uf)
	fmt.Println(uf)
	m := uf.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	r := formats.Resides{"Vancouver", "Canada"}
	pr := formats.PersonWithResides{"Ding", "Dong", r}
	fmt.Println(pr)
	q, _ := json.Marshal(pr)
	fmt.Println(q)

	var mpr formats.PersonWithResides
	json.Unmarshal(q, &mpr)
	fmt.Println(mpr)

	var uff interface{}
	json.Unmarshal(q, &uff)
	fmt.Println(uff)
	mf := uff.(map[string]interface{})
	for k, v := range mf {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case interface{}:
			fmt.Println(k, "is interface{}", vv)
			mff := vv.(map[string]interface{})
			for kk, vvv := range mff {
				switch x := vvv.(type) {
				case string:
					fmt.Println(kk, " is string ", x)
				}
			}

		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}

	}

}
