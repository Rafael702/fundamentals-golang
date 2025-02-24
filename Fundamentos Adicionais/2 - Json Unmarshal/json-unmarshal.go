package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade uint   `json: "idade"`
}

func main() {
	cachorroJSON := `{"nome":"Rex","raca":"Dalmata", "idade":3}`

    var c cachorro

    if err := json.Unmarshal([]byte(cachorroJSON), &c); err != nil {
        log.Fatal(err)
    }

    fmt.Println(c)

    cachorroJSON2 := `{"nome":"Toby","raca":"Poodle"}`

    c2 := make(map[string]string)

    if err := json.Unmarshal([]byte(cachorroJSON2), &c2); err != nil {
        log.Fatal(err)
    }

    fmt.Println(c2)
}