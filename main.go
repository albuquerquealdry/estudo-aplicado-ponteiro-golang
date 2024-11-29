package main

import "fmt"

func main() {
	var memeory_address int
	var po *int
	fmt.Println(memeory_address)
	fmt.Println(&memeory_address) // Use & para referenciar o endereço de memória.
	fmt.Println(po)               // Use & para referenciar o endereço de memória do ponteiro.

	// Entendendo como o ponteiro realmente funciona
	fmt.Println("####### ENTENDENDO COMO PONTEIROS FUNCIONAM #######")
	var personagem string
	var p *string

	personagem = "Goku"
	p = &personagem

	fmt.Println("Valor do ponteiro:" + *p) // Resgatando o valor da variavel apontada no ponteiro.
	fmt.Println(&personagem)               // Endereço de memória da variavel.
	fmt.Println(&p)                        // Endereço de memória do ponteiro.
}
