package main

import (
	"fmt"
)

//Pilha com lista ligada

type NodePilha struct {
	dado int
	prox *NodePilha
}

type Pilha struct {
	topo    *NodePilha
	tamanho int
}

func (p *Pilha) Vazia() bool {
	return p.topo == nil
}

func (p *Pilha) Empilha(dado int) {
	novoNodo := &NodePilha{}
	novoNodo.dado = dado
	novoNodo.prox = p.topo
	p.topo = novoNodo
	p.tamanho += 1
}

func (p *Pilha) Desempilha() {
	if p.Vazia() {
		fmt.Println("Pilha Vazia, não é possível desempilhar")
		return
	}

	p.topo = p.topo.prox
	p.tamanho -= 1
}

func (p *Pilha) Cabeca() (int, error) {
	if p.Vazia() {
		return 0, fmt.Errorf("Pilha Vazia")
	}
	return p.topo.dado, nil
}

func (p *Pilha) Tamanho() int {
	return p.tamanho
}

func (p *Pilha) Print() {
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

func (p *Pilha) menu() {
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
	pilha := &Pilha{}
	fmt.Println("Primeira Inicialização da Pilha")

	pilha.topo = nil
	pilha.tamanho = 0
	pilha.menu()

}
