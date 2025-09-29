package main

import (
	"flag"
	"fmt"
	"github.com/vmware/go-pmem-transaction/pmem"
	"github.com/vmware/go-pmem-transaction/transaction"
)

var (
	pmemArquivo = flag.String("file", "pilha.goPool", "nome do arquivo pmem")
)

//Pilha com lista ligada

type NodePilha struct {
	dado int
	prox *NodePilha
}

type PilhaPersistente struct {
	topo    *NodePilha
	tamanho int
}

func (p *PilhaPersistente) Vazia() bool {
	return p.topo == nil
}


func (p *PilhaPersistente) Empilha (dado int) {
			novoNodo := pnew(NodePilha)

			txn("undo") {
					novoNodo.dado = dado
					novoNodo.prox = p.topo
					p.topo = novoNodo
					p.tamanho += 1
		}
	}

	func(p *PilhaPersistente) Desempilha() {
		if p.Vazia() {
				fmt.Println("Pilha Vazia, não é possível desempilhar")
				return
		}

		txn("undo"){
			p.topo = p.topo.prox
			p.tamanho -= 1
		}
	}

	func(p *PilhaPersistente) Cabeca() (int, error) {
			if p.Vazia(){
				return 0, fmt.Errorf("Pilha Vazia")
		}
		return p.topo.dado,nil
	}

func (p *PilhaPersistente) Tamanho() int {
	return p.tamanho
}

func (p *PilhaPersistente) Print() {
	if p.Vazia() {
		fmt.Println("Pilha vazia")
	}

	fmt.Print("Pilha (topo -> base): ")
	atual := p.topo
	for atual != nil {
		fmt.Printf("%d ", atual.dado)
		atual = atual.prox
	}
}

func (p *PilhaPersistente) menu() {
	var op int
	var valor int

	for {
		fmt.Println("\n== Pilha Persistente ==")
		p.Print()
		fmt.Printf("Tamanho: %d\n", p.Tamanho())

		fmt.Println("\nEscolha uma opção:")
		fmt.Println("   1 - Empilhar")
		fmt.Println("   2 - Desempilhar")
		fmt.Println("   3 - Ver topo")
		fmt.Println("   4 - Verificar se está vazia")
		fmt.Println("   0 - Sair")

		fmt.Print("Opção: ")
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Print("Digite o valor:")
			fmt.Scan(&valor)
			p.Empilha(valor)
			fmt.Printf("Valor %d empilhado com sucesso!\n", valor)
		case 2:
			if !p.Vazia() {
				topo, _ := p.Cabeca()
				p.Desempilha()
				fmt.Printf("Valor %d desempilhado com sucesso!\n", topo)

			} else {
				fmt.Println("Pilha está vazia!")
			}
		case 3:
			if topo, err := p.Cabeca(); err != nil {
				fmt.Println("Erro:", err)
			} else {
				fmt.Printf("Topo da pilha: %d\n", topo)
			}
		case 4:
			if p.Vazia() {
				fmt.Println("A pilha está vazia")
			} else {
				fmt.Println("A pilha não está vazia")
			}
		case 0:
			fmt.Println("Saindo... Os dados estão persistindo!")
			return
		default:
			fmt.Println("Opção inválida!")

		}

	}

}

func main() {
	var pilha *PilhaPersistente 
	flag.Parse()

	primeiroInit := pmem.Init(*pmemArquivo)

	if primeiroInit {
		fmt.Println("Primeira Inicialização da Pilha")

		pilha = (*PilhaPersistente)(pmem.New("pilha",pilha))
		pilha.topo = nil 
		pilha.tamanho = 0 
	} else {
			fmt.Println("Carregando a Pilha")
			
			pilha = (*PilhaPersistente)(pmem.Get("pilha",pilha))

			fmt.Println("Pilha carregada com sucesso!")

	}
	pilha.menu()
}
