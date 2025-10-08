package main

import "testing"

func BenchmarkEmpilha(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		p := &Pilha{}
		for j := 0; j < 1000; j++ {
			p.Empilha(j)
		}
	}
}

func BenchmarkDesempilha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := &Pilha{}
		for j := 0; j < 1000; j++ {
			p.Empilha(j)
		}
		for j := 0; j < 1000; j++ {
			p.Desempilha()
		}
	}
}

func BenchmarkEmpilhaDesimpilhaAlternado(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := &Pilha{}
		for j := 0; j < 500; j++ {
			p.Empilha(j)
			p.Desempilha()
		}
	}
}

func BenchmarkCabeca(b *testing.B) {
	p := &Pilha{}
	for j := 0; j < 1000; j++ {
		p.Empilha(j)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = p.Cabeca()
	}
}
