package main

import "fmt"

// ListaGenerica do tipo interface
type ListaGenerica []interface{}

// RemoverIndice remove elemento de uma lista com base no indice informado
func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
	l := *lista
	removido := l[indice]
	*lista = append(l[0:indice], l[indice+1:]...)
	return removido
}

// RemoverInicio remove primeiro elemento de uma lista
func (lista *ListaGenerica) RemoverInicio() interface{} {
	return lista.RemoverIndice(0)
}

// RemoverFim remove último elemento de uma lista
func (lista *ListaGenerica) RemoverFim() interface{} {
	return lista.RemoverIndice(len(*lista)-1)
}

func main() {
	lista := ListaGenerica{1, "Café", 42, true, 23, "Bola", 3.14, false}

	fmt.Printf("Lista original:\n%v\n\n", lista)

	fmt.Printf("Removendo do início: %v, após remoção: \n%v\n", lista.RemoverInicio(), lista)

	fmt.Printf("Removendo do fim: %v, após remoção: \n%v\n", lista.RemoverFim(), lista)

	fmt.Printf("Removendo indice 3: %v, após remoção: \n%v\n", lista.RemoverIndice(3), lista)

	fmt.Printf("Removendo indice 0: %v, após remoção: \n%v\n", lista.RemoverIndice(0), lista)

	fmt.Printf("Removendo último indice: %v, após remoção: \n%v\n", lista.RemoverIndice(len(lista)-1), lista)
}