package main

import (
	"flag"
	"fmt"
)

// Define a estrutura do nó da lista ligada
type Node struct {
	data    int
	proximo *Node
}

// Define a estrutura da lista ligada
type ListaLigada struct {
	cabeca *Node
}

// Função para inserir um elemento no inicio da lista ligada
func (ll *ListaLigada) insereComeco(data int) {
	novoNodo := &Node{data: data}
	novoNodo.data = data
	novoNodo.proximo = ll.cabeca
	ll.cabeca = novoNodo

}

// Função para inserir um elemento no final da lista ligada
func (ll *ListaLigada) insereFim(data int) {

	novoNodo := &Node{data: data}
	novoNodo.data = data
	novoNodo.proximo = nil

	if ll.cabeca == nil {
		ll.cabeca = novoNodo
	}
	atual := ll.cabeca
	for atual.proximo != nil {
		atual = atual.proximo
	}

	atual.proximo = novoNodo
}

// Função para inserir um elemento em uma posição específica na lista ligada
func (ll *ListaLigada) inserePOS(data int, pos int) {
	novoNodo := &Node{data: data}
	novoNodo.data = data
	novoNodo.proximo = nil

	if pos == 0 {
		ll.insereComeco(data)
		return
	}
	atual := ll.cabeca
	for i := 0; i < pos-1; i++ {
		atual = atual.proximo
	}
	if atual == nil {
		fmt.Println("Posição inválida")
		return
	}
	novoNodo.proximo = atual.proximo
	atual.proximo = novoNodo
}

// Função para remover um elemento em uma posição específica da lista ligada
func (ll *ListaLigada) removePOS(pos int) {
	if pos == 0 {
		if ll.cabeca != nil {
			ll.cabeca = ll.cabeca.proximo
		} else {
			fmt.Println("Lista Vazia")
		}
		return
	}
	atual := ll.cabeca
	for i := 0; i < pos-1 && atual != nil; i++ {
		atual = atual.proximo
	}
	if atual == nil || atual.proximo == nil {
		fmt.Println("Posição inválida")
		return
	}
	atual.proximo = atual.proximo.proximo
}

// Função para imprimir elementos da lista ligada
func (ll *ListaLigada) printLista() {
	atual := ll.cabeca
	for atual != nil {
		fmt.Printf("%d -> ", atual.data)
		atual = atual.proximo
	}
	fmt.Println("nil")
}

func (ll *ListaLigada) menu() {
	var op int
	var op2 int
	var op3 int
	var info int
	op = 1
	for op != 0 {
		fmt.Println("\n Lista:")
		ll.printLista()
		fmt.Println("\nEscolha uma opção do Menu: ")
		fmt.Println("   1 - Inserir")
		fmt.Println("   2 - Remover")
		fmt.Println("   0 - Sair")

		fmt.Scan(&op)

		if op == 1 {
			fmt.Println("\nEscolha uma opção de Inserção :")
			fmt.Println("   1 - Inserir no começo")
			fmt.Println("   2 - Inserir no fim")
			fmt.Println("   3 - Inserir na posição")
			fmt.Println("   4 - Voltar ao Menu")

			fmt.Scan(&op2)

			if op2 != 4 {
				fmt.Println("\nEscolha qual o valor vai adiconar: ")
				fmt.Scan(&info)
			}
			if op2 == 1 {
				ll.insereComeco(info)
			} else if op2 == 2 {
				ll.insereFim(info)
			} else if op2 == 3 {
				fmt.Println("\nEscolha aonde quer inserir: ")
				fmt.Scan(&op3)
				ll.inserePOS(info, op3)
			}

		} else if op == 2 {
			fmt.Println("\nEscolha a posição que quer remover: ")
			fmt.Scan(&op3)
			ll.removePOS(op3)
		}
	}
}

func main() {
	flag.Parse()
	lista := &ListaLigada{}

	lista.menu()
}
