package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)


type cachorro struct {
	Nome string `json: "nome"`
	Raca string `json: "raca"`
	Idade uint `json: "idade"`
}

func main(){
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cachorroJson, err := json.Marshal(c)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroJson)
	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroJson2, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson2)
	fmt.Println(bytes.NewBuffer(cachorroJson2))

}