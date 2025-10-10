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

		insercoes := 0 
		
			for {
				valor := rand.Intn(201)
				lista.insereComeco(valor)
				insercoes++ 
				if insercoes%1000 == 0{
						fmt.Printf("Inserções até agora: %d\n",insercoes)
			}
		}
}
