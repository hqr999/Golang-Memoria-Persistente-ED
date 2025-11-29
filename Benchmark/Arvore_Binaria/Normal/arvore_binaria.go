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

type AVLNode struct {
	chave    int
	esquerda *AVLNode
	direita  *AVLNode
	altura   int
}

type ArvoreAVL struct {
	raiz *AVLNode
}

// Funções auxiliares
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pegaAltura(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return n.altura
}

func fatorBalanceamento(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return pegaAltura(n.direita) - pegaAltura(n.esquerda)
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

// Inserção, deleção e busca
func insere_tmpl(raiz *AVLNode, chave int) *AVLNode {
	if raiz == nil {
		return &AVLNode{chave: chave, altura: 1}
	}

	if chave < raiz.chave {
		raiz.esquerda = insere_tmpl(raiz.esquerda, chave)
	} else if chave > raiz.chave {
		raiz.direita = insere_tmpl(raiz.direita, chave)
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

func valorMinNodo(n *AVLNode) *AVLNode {
	atual := n
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
				raiz = nil
			} else {
				*raiz = *temp
			}
		} else {
			temp := valorMinNodo(raiz.direita)
			raiz.chave = temp.chave
			raiz.direita = deletaNodo_tmpl(raiz.direita, temp.chave)
		}
	}

	if raiz == nil {
		return nil
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

func buscar(raiz *AVLNode, chave int) *AVLNode {
	if raiz == nil || raiz.chave == chave {
		return raiz
	}

	if chave < raiz.chave {
		return buscar(raiz.esquerda, chave)
	}
	return buscar(raiz.direita, chave)
}

// Métodos da árvore
func (a *ArvoreAVL) inserir(chave int) {
	a.raiz = insere_tmpl(a.raiz, chave)
}

func (a *ArvoreAVL) deletar(chave int) {
	a.raiz = deletaNodo_tmpl(a.raiz, chave)
}

func (a *ArvoreAVL) buscarValor(chave int) *AVLNode {
	return buscar(a.raiz, chave)
}

// Workloads
func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	arvore := &ArvoreAVL{}

	fmt.Printf("🚀 Executando workload: %s\n", *workload)

	inserts, updates, deletes := 0, 0, 0

	switch *workload {
	case "insert":
		for {
			arvore.inserir(rand.Intn(200))
			inserts++
			if inserts%1000 == 0 {
				fmt.Printf("Inserções até agora: %d\n", inserts)
			}
		}

	case "read":
		for {
			if rand.Float64() < 0.5 {
				arvore.inserir(rand.Intn(200))
				inserts++
			} else {
				n := arvore.buscarValor(rand.Intn(200))
				if n != nil {
					n.chave = rand.Intn(200)
					updates++
				}
			}
			if (inserts+updates)%1000 == 0 {
				fmt.Printf("Ops: %d (insert: %d | update: %d)\n", inserts+updates, inserts, updates)
			}
		}

	case "delete":
		for {
			if rand.Float64() < 0.5 {
				arvore.inserir(rand.Intn(200))
				inserts++
			} else {
				arvore.deletar(rand.Intn(200))
				deletes++
			}
			if (inserts+deletes)%1000 == 0 {
				fmt.Printf("Ops: %d (insert: %d | delete: %d)\n", inserts+deletes, inserts, deletes)
			}
		}

	default:
		fmt.Println("Workload inválido. Use -workload=insert | update | delete")
	}

}
