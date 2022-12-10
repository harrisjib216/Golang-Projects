package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Res1 struct {
	Page int
	Fruits []string
}

type Res2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

var p = fmt.Println

func main() {
	boolB, _ := json.Marshal(true)
	p(string(boolB))

	intB, _ := json.Marshal(1)
	p(string(intB))

	floatB, _ := json.Marshal(22.34)
	p(string(floatB))

	strB, _ := json.Marshal("golang")
	p(string(strB))


	sliceD := []string{"apple", "peach", "pear"}


	sliceB, _ := json.Marshal(sliceD)
	p(sliceB)

	mapB, _ := json.Marshal(map[string]int{"applice": 5, "lettuce": 7})
	p(mapB)

	res1D := &Res1{
		Page: 1,
		Fruits: sliceD,
	}
	res1B, _ := json.Marshal(res1D)
	p(string(res1B))

	res2D := &Res2{
		Page: 1,
		Fruits: sliceD,
	}
	res2B, _ := json.Marshal(res2D)
	p(res2B)


	byteD := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)
	
	var body map[string]interface{}

	if err := json.Unmarshal(byteD, &body); err != nil {
        panic(err)
    }
	p(body)

	num := body["num"].(float64)
	p(num)

	strs := body["strs"].([]interface{})
	str1 := strs[0].(string)
	p(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Res2{}
	json.Unmarshal([]byte(str), &res)
	p(res)
	p(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}

	enc.Encode(d)
}