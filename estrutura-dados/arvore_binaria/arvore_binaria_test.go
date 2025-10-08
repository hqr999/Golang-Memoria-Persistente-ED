package main

import (
	"math/rand"
	"testing"
	"time"
)

// Benchmark para inserção sequencial
func BenchmarkInserirSequencial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arv := &ArvoreAVL{}
		for j := 0; j < 1000; j++ {
			arv.inserir(j)
		}
	}
}

// Benchmark de inserção aleatória
func BenchmarkInserirAleatoria(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		arv := &ArvoreAVL{}
		for j := 0; j < 1000; j++ {
			arv.inserir(rand.Intn(100000))
		}
	}
}

// Benchmark de deleção aleatória
func BenchmarkDeletarAleatoria(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		arv := &ArvoreAVL{}
		val := rand.Perm(1000)
		for _, v := range val {
			arv.inserir(v)
		}
		for _, v := range val {
			arv.deletar(v)
		}
	}
}

// Benchmark de deleção sequencial
func BenchmarkDeletarSequencial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arv := &ArvoreAVL{}
		for j := 0; j < 1000; j++ {
			arv.inserir(j)
		}
		for j := 0; j < 1000; j++ {
			arv.deletar(j)
		}
	}
}
