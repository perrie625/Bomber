package main

import (
	"encoding/json"
	"io/ioutil"
)

type test struct {
	A string `json:"b"`
	B string `json:"a"`
}

func main() {
	file, _ := ioutil.ReadFile("./test.json")
	tmp := test{}
	json.Unmarshal(file, &tmp)
	println(tmp.A)
	println(tmp.B)
}
