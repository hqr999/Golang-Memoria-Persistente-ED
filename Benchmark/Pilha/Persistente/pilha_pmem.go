package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/vmware/go-pmem-transaction/pmem"
	"github.com/vmware/go-pmem-transaction/transaction"
)

var (
	pmemArquivo = flag.String("file", "pilha.goPool", "nome do arquivo pmem")
	workload = flag.String("workload","insert","tipo de workload: insert | update | delete")
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

	func(p *PilhaPersistente) Cabeca() (int, bool) {
			if p.Vazia(){
				return 0, false 
		}
		return p.topo.dado,true 
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



func main() {
	var pilha *PilhaPersistente 
	flag.Parse()

	primeiroInit := pmem.Init(*pmemArquivo)

	if primeiroInit {

		pilha = (*PilhaPersistente)(pmem.New("pilha",pilha))
		pilha.topo = nil 
		pilha.tamanho = 0 
	} else {
			pilha = (*PilhaPersistente)(pmem.Get("pilha",pilha))
	}
	fmt.Printf("🚀 Executando workload: %s\n", *workload)

	pushes, pops, reads := 0, 0, 0

	switch *workload{
			case "insert":
						for {
							pilha.Empilha(rand.Intn(200))
							pushes++
							if pushes%1000 == 0{
								fmt.Printf("Empilhamentos: %d\n",pushes)		
							}
						}
			case "update":
					for {
							if rand.Float64() < 0.5{
									pilha.Empilha(rand.Intn(200))
									pushes++
							} else{
									_ , ok := pilha.Cabeca()
									if ok{
												reads++
									}
							}
							if (pushes+reads)%1000 == 0{
									fmt.Printf("Ops: %d (push: %d | read: %d)\n",pushes+reads,pushes,reads)
							}
					}
			case "delete":
				for {
						if rand.Float64() < 0.5 {
							pilha.Empilha(rand.Intn(200))
							pushes++ 
						}	else if !pilha.Vazia(){
									pilha.Desempilha()
									pops++
							}
						if (pushes+pops)%1000 == 0 {
							fmt.Printf("Ops: %d(push: %d | pop: %d | tamanho: %d)",pushes+pops,pushes,pops,pilha.tamanho)
						}
				}
			default:
					fmt.Println("Workload inválido. Use -workload=inser | update | delete")
		}
}
