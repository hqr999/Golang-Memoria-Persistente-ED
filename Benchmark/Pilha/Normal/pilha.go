package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	workload = flag.String("workload", "insert", "tipo de workload: insert | read | delete")
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

func (p *Pilha) Cabeca() (int, bool) {
	if p.Vazia() {
		return 0, false
	}
	return p.topo.dado, true
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

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("🚀 Executando workload: %s\n", *workload)

	pilha := &Pilha{}
	pushes, pops, reads := 0, 0, 0

	switch *workload {
	case "insert":
		for {
			pilha.Empilha(rand.Intn(200))
			pushes++
			if pushes%1000 == 0 {
				fmt.Printf("Empilhamentos: %d\n", pushes)
			}
		}
	case "read":
		for {
			if rand.Float64() < 0.5 {
				pilha.Empilha(rand.Intn(200))
				pushes++
			} else {
				_, ok := pilha.Cabeca()
				if ok {
					reads++
				}
			}
			if (pushes+reads)%1000 == 0 {
				fmt.Printf("Ops: %d (push: %d | read: %d)\n", pushes+reads, pushes, reads)
			}
		}
	case "delete":
		for {
			if rand.Float64() < 0.5 {
				pilha.Empilha(rand.Intn(200))
				pushes++
			} else if !pilha.Vazia() {
				pilha.Desempilha()
				pops++
			}
			if (pushes+pops)%1000 == 0 {
				fmt.Printf("Ops: %d(push: %d | pop: %d | tamanho: %d)", pushes+pops, pushes, pops, pilha.tamanho)
			}
		}
	default:
		fmt.Println("Workload inválido. Use -workload=insert | update | delete")
	}
}
