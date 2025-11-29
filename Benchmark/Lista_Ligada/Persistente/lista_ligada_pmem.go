package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/vmware/go-pmem-transaction/transaction"
	"github.com/vmware/go-pmem-transaction/pmem"

)

var(
		arquivoPmem = flag.String("file","lista-ligada.goPool","arquivo de memória persistente")
	workload = flag.String("workload","insert","tipo do workload: insert | read | delete")
)

type Node struct {
	data int 
	prox *Node 
}

type ListaLigada struct{
		cabeca *Node 
		cauda *Node 
		tam int 
}

func (ll *ListaLigada) insereComeco(data int){
		novoNodo := pnew(Node)
		
		txn("undo") {
			novoNodo.data = data 
			novoNodo.prox = ll.cabeca 
			ll.cabeca = novoNodo 
			if ll.tam == 0 {
					ll.cauda = novoNodo
			}
		ll.tam++
		}
}

func(ll *ListaLigada) atualizaAleatorio(){
		if ll.tam == 0 {
			return
		}
	pos := rand.Intn(ll.tam)
	atual := ll.cabeca
	for i := 0; i < pos && atual.prox != nil ; i++ {
		atual = atual.prox
	}
	txn("undo"){
		atual.data = rand.Intn(201)
	}
}


func(ll *ListaLigada) removePOS(pos int){
		if pos < 0 || pos >= ll.tam{ 
			return 
		}

		txn("undo"){
			if pos == 0 {
				ll.cabeca = ll.cabeca.prox 
				if ll.tam == 1 {
					ll.cauda = nil
				}
				ll.tam-- 
				return 
			}
			anterior := ll.cabeca 
			for i := 0; i < pos-1 && anterior.prox != nil; i++ {
					anterior = anterior.prox
			}

			removido := anterior.prox 
			if removido == nil {
				return 
			}
			anterior.prox = removido.prox 
			if pos == ll.tam - 1 {
				ll.cauda = anterior
			}
			ll.tam--
	}
}

func main() {
		flag.Parse()
		rand.Seed(time.Now().UnixNano())
	
		var lista *ListaLigada
		primeiraInit := pmem.Init(*arquivoPmem)

		if primeiraInit{
			lista = (*ListaLigada)(pmem.New("root",lista))
			lista.cabeca = nil 
			lista.cauda = nil 
			lista.tam = 0
			fmt.Println("Primeira inicialização do pool PMEM")
		} else{
				lista = (*ListaLigada)(pmem.Get("root",lista))
		
		}
		fmt.Printf("Executando workload: %s\n",*workload)
		insercoes := 0 
		atualizacoes := 0
		delecoes := 0

		switch *workload{
			case "insert":
				for {
					valor := rand.Intn(201)
					lista.insereComeco(valor)
					insercoes++ 
					if insercoes%1000 == 0{
						fmt.Printf("Inserções até agora: %d\n",insercoes)
					}
				}
			case "read":
					for {
							if rand.Float64() < 0.5 {
									lista.insereComeco(rand.Intn(201)) //50% inserção 								 
									insercoes++
							} else {
									lista.atualizaAleatorio() //50% atualização 
									atualizacoes++ 
							}
						
							if (insercoes+atualizacoes)%1000 == 0{
									fmt.Printf("Ops: %d (insert: %d | update: %d)\n",insercoes+atualizacoes,insercoes,atualizacoes)
						}
					}
			case "delete":
					for {
								if rand.Float64() < 0.5 {
									lista.insereComeco(rand.Intn(201)) //50% inserção 								 
									insercoes++
							} else if lista.tam > 0 {
								pos := rand.Intn(lista.tam)
								lista.removePOS(pos)
								delecoes++
							}
						
							if (insercoes+delecoes)%1000 == 0{
				fmt.Printf("Ops: %d (insert: %d | delete: %d | tam: %d)\n",insercoes+delecoes,insercoes,delecoes, lista.tam)
						}

					}
			default:
					fmt.Println("Workload inválido. Use -workload=insert ou -workload=update")
		}	
}
