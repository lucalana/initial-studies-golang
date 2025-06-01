package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// No GO após campo de uma struct fazemos dessa forma para mostrar como o campo vai ser mostado
// quando convertido para json
type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
}

func main() {
	// CRIANDO UM JSON
	cao1 := cachorro{"Jose", "Pitbull"}
	cao2 := map[string]string{
		"nome": "Jose",
		"raca": "Pitbull",
	}

	fmt.Println(cao1)
	fmt.Println(cao2)
	cao1Json, err := json.Marshal(cao1)
	if err != nil {
		fmt.Println(err)
	}
	cao2Json, err := json.Marshal(cao1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.NewBuffer(cao2Json))
	fmt.Println(bytes.NewBuffer(cao1Json))

	// PASSANDO UM JSON PARA UM STRUCT/MAP
	caoJson := `{"nome":"Jose", "raca":"Pitbull"}`
	var c cachorro
	c2 := make(map[string]string)
	fmt.Println(caoJson)

	if err := json.Unmarshal([]byte(caoJson), &c); err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal([]byte(caoJson), &c2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
	fmt.Println(c2)
}
