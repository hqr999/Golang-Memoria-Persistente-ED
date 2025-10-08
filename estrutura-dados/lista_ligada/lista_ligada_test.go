package main

import "testing"

func criaLista(n int) *ListaLigada {
	ll := &ListaLigada{}
	for i := 0; i < n; i++ {
		ll.insereFim(i)
	}
	return ll
}

// Benchmark para inserção no começo
func BenchmarkInsereComeco(b *testing.B) {
	ll := &ListaLigada{}
	for i := 0; i < b.N; i++ {
		ll.insereComeco(i)
	}
}

// Benchmark para inserção no fim
func BenchmarkInsereFim(b *testing.B) {
	ll := &ListaLigada{}
	for i := 0; i < b.N; i++ {
		ll.insereFim(i)
	}
}

// Benchmark para inserção em posição específica
func BenchmarkInserePOS(b *testing.B) {
	//primeiro preenchemos a lista com b.N elementos
	for i := 0; i < b.N; i++ {
		ll := criaLista(1000)
		ll.inserePOS(9999, 500)
	}

}

// Benchmark para remoção em posição específica
func BenchmarkRemovePOS(b *testing.B) {

	//preenchemos a lista
	for i := 0; i < b.N; i++ {
		ll := criaLista(1000)
		ll.removePOS(500)
	}
}
