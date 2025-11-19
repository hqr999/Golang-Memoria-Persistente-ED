package main

import (
	"os"
	"flag"
	"math/rand"
	"time"
	"fmt"

	"github.com/vmware/go-pmem-transaction/transaction"

	"github.com/vmware/go-pmem-transaction/pmem"
)

var (
	arquivoPmem = flag.String("file", "lista.goPool", "arquivo de memória persistente")
)


type Node struct {
	data int
	prox  *Node
}

type ListaLigada struct{
	cabeca *Node 
	cauda *Node 
	tam int 
}

func(ll *ListaLigada) insereInicio(valor int){
	n := pnew(Node)

	txn("undo"){
		n.data = valor 
		n.prox = ll.cabeca 
		ll.cabeca = n 
		
		if ll.tam == 0{
			ll.cauda = n
		}

		ll.tam++
	}
}

func main(){
		flag.Parse() 
		rand.Seed(time.Now().UnixNano())

		var lista *ListaLigada 
		primeira := pmem.Init(*arquivoPmem)
		
		if primeira{
				fmt.Println("Primeira inicialização do pool.")
				lista = (*ListaLigada)(pmem.New("root", lista))

				txn("undo") {
					lista.cabeca = nil
					lista.cauda = nil
					lista.tam = 0
				}	
		} else{
			lista = (*ListaLigada)(pmem.Get("root", lista))
			fmt.Printf("Pool existente. Lista contém %d elementos.\n", lista.tam)
	}
		fmt.Printf("Processo iniciado com o nome de: %v \n",os.Args[0])
		fmt.Println("Inserindo valores infinitamente. Mate o processo para testar crash-consistency.")

		contador := 0
		for {
			val := rand.Intn(1000)
			lista.insereInicio(val)
			
			contador++

			if contador%1000 == 0 {
				fmt.Printf("Inseridos: %d (tamanho atual: %d)\n", contador, lista.tam)

				time.Sleep(1 * time.Second)
		}
	}
}
