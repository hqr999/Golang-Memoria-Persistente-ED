package main

import (
	"flag"
	"fmt"
   "github.com/vmware/go-pmem-transaction/transaction"
	"github.com/vmware/go-pmem-transaction/pmem"
)

var (
	pmemArquivo = flag.String("file", "arvore_binaria.goPool", "nome do arquivo pmem")
)

type AVLNode struct {
	chave    int
	esquerda *AVLNode
	direita  *AVLNode
	altura   int
}

func max(a, b int) int {
	switch {
	case a > b:
		return a
	default:
		return b

	}
}

func pegaAltura(nodo *AVLNode) int {
	if nodo == nil {
		return 0
	}
	return nodo.altura
}

func fatorBalanceamento(nodo *AVLNode) int {
	switch {
	case nodo == nil:
		return 0
	default:
		return pegaAltura(nodo.direita) - pegaAltura(nodo.esquerda)
	}
}

func rotacionarDireita(y *AVLNode) *AVLNode {
	x := y.esquerda
	T2 := x.direita

	x.direita = y
	y.esquerda = T2

	y.altura = max(pegaAltura(y.esquerda), pegaAltura(y.direita)) + 1
	x.altura = max(pegaAltura(x.esquerda), pegaAltura(x.direita)) + 1

	return x
}

func rotacionarEsquerda(x *AVLNode) *AVLNode {
	y := x.direita
	T2 := y.esquerda

	y.esquerda = x
	x.direita = T2

	x.altura = max(pegaAltura(x.esquerda), pegaAltura(x.direita)) + 1
	y.altura = max(pegaAltura(y.esquerda), pegaAltura(y.direita)) + 1

	return y
}

func insere_tmpl(raiz *AVLNode, chave int) *AVLNode {
	if raiz == nil {
		novoNodo := pnew(AVLNode)

		novoNodo.chave = chave
		novoNodo.altura = 1
		novoNodo.direita = nil
		novoNodo.esquerda = nil

		return novoNodo
	}

	if chave < raiz.chave {
		raiz.esquerda = insere_tmpl(raiz.esquerda,chave)
	} else if chave > raiz.chave {
		raiz.direita = insere_tmpl(raiz.direita,chave)
	} else {
		return raiz
	}

	raiz.altura = 1 + max(pegaAltura(raiz.esquerda), pegaAltura(raiz.direita))
	bf := fatorBalanceamento(raiz)

	if bf < -1 {
		if fatorBalanceamento(raiz.esquerda) <= 0 {
			return rotacionarDireita(raiz)
		} else {
			raiz.esquerda = rotacionarEsquerda(raiz.esquerda)
			return rotacionarDireita(raiz)
		}
	} else if bf > 1 {
		if fatorBalanceamento(raiz.direita) >= 0 {
			return rotacionarEsquerda(raiz)
		} else {
			raiz.direita = rotacionarDireita(raiz.direita)
			return rotacionarEsquerda(raiz)
		}
	}
	return raiz
}

func valorMinNodo(node *AVLNode) *AVLNode{
		atual := node 
		for atual.esquerda != nil {
			atual = atual.esquerda
	}
	return atual 
}

func deletaNodo_tmpl(raiz *AVLNode,chave int) *AVLNode {
		if raiz == nil {
			return nil 
	}

	if chave < raiz.chave {
		raiz.esquerda = deletaNodo_tmpl(raiz.esquerda,chave)
	} else if chave > raiz.chave {
		raiz.direita = deletaNodo_tmpl(raiz.direita,chave)
	} else {
			if raiz.esquerda == nil || raiz.direita == nil {
				var temp *AVLNode 
				if raiz.esquerda != nil {
					temp = raiz.esquerda
			} else {
				temp = raiz.direita
			}
			
			if temp == nil {
					temp = raiz 
					raiz = nil 
			} else {
				*raiz = *temp 
			}
		}else {
			temp := valorMinNodo(raiz.direita)
			raiz.chave = temp.chave
			raiz.direita = deletaNodo_tmpl(raiz.direita,temp.chave)
		}

	}
	if raiz == nil {
			return nil 
	}
	raiz.altura = 1 + max(pegaAltura(raiz.esquerda),pegaAltura(raiz.direita))
	bf := fatorBalanceamento(raiz)
	if bf < -1 {
			if fatorBalanceamento(raiz.esquerda) <= 0 {
				return rotacionarDireita(raiz)
		} else {
				raiz.esquerda = rotacionarEsquerda(raiz.esquerda)
				return rotacionarDireita(raiz)
		}
	} else if bf > 1 {
		if fatorBalanceamento(raiz.direita) >= 0 {
				return rotacionarEsquerda(raiz)
		} else {
				raiz.direita = rotacionarDireita(raiz.direita)
				return rotacionarEsquerda(raiz)
		} 
	}

	return raiz
}


func buscar(raiz *AVLNode, chave int) *AVLNode {
	if raiz == nil || raiz.chave == chave {
			return raiz 
	}	
	if chave < raiz.chave {
		return buscar(raiz.esquerda,chave)
	}
	return buscar(raiz.direita,chave)
}

func inorder(raiz *AVLNode) {
		if raiz != nil {
		inorder(raiz.esquerda)
		fmt.Printf("%d ",raiz.chave)
		inorder(raiz.direita)
	}
}

func main() {
	
	var raiz *AVLNode 

	flag.Parse()
	primeiroInit := pmem.Init(*pmemArquivo)

	if primeiroInit {
		raiz = (*AVLNode)(pmem.New("root",raiz))
		
		txn("undo"){
			raiz = insere_tmpl(raiz,11)
			raiz = insere_tmpl(raiz,22)
			raiz = insere_tmpl(raiz,55)
			raiz = insere_tmpl(raiz,44)
			raiz = insere_tmpl(raiz,77)
			raiz = insere_tmpl(raiz,66)
		}

		fmt.Print("Inorder traversal: ")
		inorder(raiz)
		fmt.Println()

	} else {
			raiz = (*AVLNode)(pmem.Get("root",raiz))
			inorder(raiz)
	}

}
