package main

import (
	"flag"
	"fmt"
"github.com/vmware/go-pmem-transaction/pmem"
	"github.com/vmware/go-pmem-transaction/transaction"
)

var (
	arquivoPmem = flag.String("file","lista-ligada.goPool","nome do arquivo pmem")
)

// Define a estrutura do nó da lista ligada 
type Node struct {
		data int 
		proximo *Node 
}

//Define a estrutura da lista ligada 
type ListaLigada struct {
		cabeca *Node
		cauda *Node 
		tam int 
}

// Função para inserir um elemento no inicio da lista ligada 
func (ll *ListaLigada) insereComeco(data int){
	novoNodo := pnew(Node)

	txn("undo") {
			novoNodo.data = data
			novoNodo.proximo = ll.cabeca
			ll.cabeca = novoNodo
			if ll.tam == 0 {
				ll.cauda = novoNodo
		}
		ll.tam++
	}

}

// Função para inserir um elemento no final da lista ligada 
func (ll *ListaLigada) insereFim(data int) {

			novoNodo := pnew(Node)
			novoNodo.data = data 
			novoNodo.proximo = nil 

				txn("undo") {
						if ll.tam == 0 {
							ll.cabeca = novoNodo 
							ll.cauda = novoNodo
						} else {
							ll.cauda.proximo = novoNodo 
							ll.cauda = novoNodo
						}
						ll.tam++
		}
}

// Função para inserir um elemento em uma posição específica na lista ligada 
func(ll *ListaLigada) inserePOS(data int,pos int) {
		if pos < 0 || pos > ll.tam {
			fmt.Println("Posição inválida")
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

		novoNodo := pnew(Node)
		novoNodo.data = data 

		atual := ll.cabeca
		for i := 0; i < pos-1; i++ {
			atual = atual.proximo
		}

		txn("undo"){
			novoNodo.proximo = atual.proximo 
			atual.proximo = novoNodo 
			ll.tam++
		}
}

//Função para remover um elemento em uma posição específica da lista ligada
func (ll *ListaLigada) removePOS(pos int) {
	if pos < 0 || pos >= ll.tam {
		fmt.Println("Posição inválida")
		return
	}

	if pos == 0 {
		txn("undo"){
			ll.cabeca = ll.cabeca.proximo
			if ll.tam == 1{
				ll.cauda = nil
			}
			ll.tam--
		}
		return
	}

	atual := ll.cabeca 
	for i := 0; i < pos-1; i++ {
		atual = atual.proximo
	}

	txn("undo"){
		removido := atual.proximo 
		atual.proximo = removido.proximo 
		if pos == ll.tam - 1{
			ll.cauda = atual
		}
		ll.tam--
	}
}

//Função para imprimir elementos da lista ligada 
func (ll *ListaLigada) printLista() {
		atual := ll.cabeca
		for atual != nil {
			fmt.Printf("%d -> ",atual.data)
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
		for (op != 0){
			fmt.Println("\n Lista:")
			ll.printLista()
			fmt.Println("\nEscolha uma opção do Menu: ")
			fmt.Println("   1 - Inserir")
			fmt.Println("   2 - Remover")
			fmt.Println("   0 - Sair")
			
			fmt.Scan(&op)
			
			if(op == 1){
					fmt.Println("\nEscolha uma opção de Inserção :")
          fmt.Println("   1 - Inserir no começo")
          fmt.Println("   2 - Inserir no fim")
          fmt.Println("   3 - Inserir na posição")
          fmt.Println("   4 - Voltar ao Menu")

					fmt.Scan(&op2)
					
			if(op2 != 4){
				fmt.Println("\nEscolha qual o valor vai adiconar: ")
						fmt.Scan(&info)
			}
			if (op2 == 1){
					ll.insereComeco(info)
			} else if (op2 == 2) {
					ll.insereFim(info)
			} else if(op2 == 3) {
					fmt.Println("\nEscolha aonde quer inserir: ")
					fmt.Scan(&op3)
					ll.inserePOS(info,op3)
			}

		} else if (op == 2) {
				fmt.Println("\nEscolha um elemento que quer remover: ")
			fmt.Scan(&op3)
			ll.removePOS(op3)
		}
	}
}

func main() {
		var lista *ListaLigada
		flag.Parse() 
		primeiroInit := pmem.Init(*arquivoPmem)

		if primeiroInit {
			//Cria um lista ligada vazia 
			lista := (*ListaLigada)(pmem.New("root",lista))
			lista.cabeca = nil 
			
			lista.menu()
	} else {
			//Não é uma primeira inicialização 
			lista = (*ListaLigada)(pmem.Get("root",lista))
			lista.menu()
	}
}
