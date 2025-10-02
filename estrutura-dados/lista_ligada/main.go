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
	cauda  *Node
	tam    int
}

// Função para inserir um elemento no inicio da lista ligada
func (ll *ListaLigada) insereComeco(data int) {
	novoNodo := &Node{data: data, proximo: ll.cabeca}
	ll.cabeca = novoNodo

	if ll.tam == 0 {
		ll.cauda = novoNodo
	}

	ll.tam++

}

// Função para inserir um elemento no final da lista ligada
func (ll *ListaLigada) insereFim(data int) {

	novoNodo := &Node{data: data}

	if ll.tam == 0 {
		ll.cabeca = novoNodo
		ll.cauda = novoNodo
	} else {
		ll.cauda.proximo = novoNodo
		ll.cauda = novoNodo
	}
	ll.tam++

}

// Função para inserir um elemento em uma posição específica na lista ligada
func (ll *ListaLigada) inserePOS(data int, pos int) {

	if pos < 0 || pos > ll.tam {
		fmt.Println("Posição Inválida")
		return
	}

	if pos == 0 {
		ll.insereComeco(data)
		return
	}
	if pos == ll.tam {
		ll.insereFim(data)
		return
	}

	atual := ll.cabeca
	for i := 0; i < pos-1; i++ {
		atual = atual.proximo
	}
	novoNodo := &Node{data: data, proximo: atual.proximo}
	atual.proximo = novoNodo
	ll.tam++

}

// Função para remover um elemento em uma posição específica da lista ligada
func (ll *ListaLigada) removePOS(pos int) {

	if pos < 0 || pos >= ll.tam {
		fmt.Println("Posição inválida")
		return
	}

	if pos == 0 {
		//Remover do começo
		ll.cabeca = ll.cabeca.proximo
		if ll.tam == 1 {
			//lista ficou vazia
			ll.cauda = nil
		}
		ll.tam--
		return
	}
	atual := ll.cabeca
	for i := 0; i < pos-1; i++ {
		atual = atual.proximo
	}

	removido := atual.proximo
	atual.proximo = removido.proximo

	if pos == ll.tam-1 {
		// se removermos o último elemento, atualizar a cauda
		ll.cauda = atual
	}
	ll.tam--

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
