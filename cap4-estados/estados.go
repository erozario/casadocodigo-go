package main

import (
	"fmt"
)

// Estado ...
type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	estados := make(map[string]Estado, 3)

	estados["GO"] = Estado{"Goiás", 6434052, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 3914419, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}

	fmt.Println(estados)
}
