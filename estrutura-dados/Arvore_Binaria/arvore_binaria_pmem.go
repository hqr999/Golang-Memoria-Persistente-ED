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

// Estrutura para manter uma referência à raiz da árvore
type ArvoreAVL struct {
	raiz *AVLNode
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
		
		txn("undo") {
			novoNodo.chave = chave
			novoNodo.altura = 1
			novoNodo.direita = nil
			novoNodo.esquerda = nil
		}
		
		return novoNodo
	}

	if chave < raiz.chave {
		raiz.esquerda = insere_tmpl(raiz.esquerda, chave)
	} else if chave > raiz.chave {
		raiz.direita = insere_tmpl(raiz.direita, chave)
	} else {
		return raiz
	}

	txn("undo") {
		raiz.altura = 1 + max(pegaAltura(raiz.esquerda), pegaAltura(raiz.direita))
	}
	
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

func valorMinNodo(node *AVLNode) *AVLNode {
	atual := node
	for atual.esquerda != nil {
		atual = atual.esquerda
	}
	return atual
}

func deletaNodo_tmpl(raiz *AVLNode, chave int) *AVLNode {
	if raiz == nil {
		return nil
	}

	if chave < raiz.chave {
		raiz.esquerda = deletaNodo_tmpl(raiz.esquerda, chave)
	} else if chave > raiz.chave {
		raiz.direita = deletaNodo_tmpl(raiz.direita, chave)
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
				txn("undo") {
					*raiz = *temp
				}
			}
		} else {
			temp := valorMinNodo(raiz.direita)
			txn("undo") {
				raiz.chave = temp.chave
			}
			raiz.direita = deletaNodo_tmpl(raiz.direita, temp.chave)
		}
	}

	if raiz == nil {
		return nil
	}

	txn("undo") {
		raiz.altura = 1 + max(pegaAltura(raiz.esquerda), pegaAltura(raiz.direita))
	}
	
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
		return buscar(raiz.esquerda, chave)
	}
	return buscar(raiz.direita, chave)
}

func inorder(raiz *AVLNode) {
	if raiz != nil {
		inorder(raiz.esquerda)
		fmt.Printf("%d ", raiz.chave)
		inorder(raiz.direita)
	}
}

func (arvore *ArvoreAVL) inserir(chave int) {
	txn("undo") {
		arvore.raiz = insere_tmpl(arvore.raiz, chave)
	}
}

func (arvore *ArvoreAVL) deletar(chave int) {
	txn("undo") {
		arvore.raiz = deletaNodo_tmpl(arvore.raiz, chave)
	}
}

func (arvore *ArvoreAVL) buscarValor(chave int) *AVLNode {
	return buscar(arvore.raiz, chave)
}

func (arvore *ArvoreAVL) imprimir() {
	fmt.Print("Árvore AVL (inorder): ")
	inorder(arvore.raiz)
	fmt.Println()
}

func main() {
	var arvore *ArvoreAVL

	flag.Parse()
	primeiroInit := pmem.Init(*pmemArquivo)

	if primeiroInit {
		fmt.Println("First time initialization")
		
		// Criar árvore AVL persistente
		arvore = (*ArvoreAVL)(pmem.New("root", arvore))
		arvore.raiz = nil

		// Inserir valores iniciais
		arvore.inserir(11)
		arvore.inserir(22)
		arvore.inserir(55)
		arvore.inserir(44)
		arvore.inserir(77)
		arvore.inserir(66)

		fmt.Print("Inorder traversal (primeira execução): ")
		arvore.imprimir()

	} else {
		fmt.Println("Loading existing data")
		// Carregar árvore AVL persistente
		arvore = (*ArvoreAVL)(pmem.Get("root", arvore))
		
		fmt.Print("Inorder traversal (dados carregados): ")
		arvore.imprimir()
		arvore.deletar(44)
		fmt.Print("Inorder traversal (dados carregados) depois do nodo 44 ser deletado: ")
		arvore.imprimir()
	}
}
