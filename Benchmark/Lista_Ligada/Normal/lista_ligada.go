package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	workload = flag.String("workload", "insert", "tipo do workload: insert | update")
)

type Node struct {
	data int
	prox *Node
}

type ListaLigada struct {
	cabeca *Node
	cauda  *Node
	tam    int
}

func (ll *ListaLigada) insereComeco(data int) {
	novoNodo := &Node{}

	novoNodo.data = data
	novoNodo.prox = ll.cabeca
	ll.cabeca = novoNodo
	if ll.tam == 0 {
		ll.cauda = novoNodo
	}
	ll.tam++

}

func (ll *ListaLigada) atualizaAleatorio() {
	if ll.tam == 0 {
		return
	}
	pos := rand.Intn(ll.tam)
	atual := ll.cabeca
	for i := 0; i < pos && atual.prox != nil; i++ {
		atual = atual.prox
	}
	atual.data = rand.Intn(201)

}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	lista := &ListaLigada{}

	lista.cabeca = nil
	lista.cauda = nil
	lista.tam = 0

	fmt.Printf("Executando workload: %s\n", *workload)
	insercoes := 0
	atualizacoes := 0

	switch *workload {
	case "insert":
		for {
			valor := rand.Intn(201)
			lista.insereComeco(valor)
			insercoes++
			if insercoes%1000 == 0 {
				fmt.Printf("Inserções até agora: %d\n", insercoes)
			}
		}
	case "update":
		for {
			if rand.Float64() < 0.5 {
				lista.insereComeco(rand.Intn(201)) //50% inserção
				insercoes++
			} else {
				lista.atualizaAleatorio() //50% atualização
				atualizacoes++
			}

			if (insercoes+atualizacoes)%1000 == 0 {
				fmt.Printf("Ops: %d (insert: %d | update: %d)\n", insercoes+atualizacoes, insercoes, atualizacoes)
			}
		}
	default:
		fmt.Println("Workload inválido. Use -workload=insert ou -workload=update")
	}
}
