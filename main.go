package main

import (
	"encoding/json"
	"fmt"

	"github.com/icebattle/jsonpractice/formats"
)

func main() {
	person := formats.Person{"Happy", "Golucky"}
	fmt.Println(person)
	buff, _ := json.Marshal(person)
	fmt.Println(buff)

	var mperson formats.Person
	json.Unmarshal(buff, &mperson)
	fmt.Println(mperson)

	var uf interface{}
	json.Unmarshal(buff, &uf)
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
	pr := formats.PersonWithResides{"Holly", "Golightly", r}
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
