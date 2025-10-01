package main

import "testing"

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
	ll := &ListaLigada{}

	//primeiro preenchemos a lista com b.N elementos
	for i := 0; i < b.N; i++ {
		ll.insereFim(i)
	}

	//agora medimos apenas inserções no meio
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.inserePOS(99999, i/2)
	}
}

// Benchmark para remoção em posição específica
func BenchmarkRemovePOS(b *testing.B) {
	ll := &ListaLigada{}

	//preenchemos a lista
	for i := 0; i < b.N; i++ {
		ll.insereFim(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.removePOS(0) // primeiro a sair
	}
}
