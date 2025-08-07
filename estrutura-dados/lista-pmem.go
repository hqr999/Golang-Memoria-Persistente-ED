package estruturadados

import (
	"fmt"
	"math/rand"

	"github.com/vmware/go-pmem-transaction/transaction"
)


const Magic = 0xCAFEBABE12345678

//Raiz é a estrutura persistente principal
type Raiz struct {
	Magic int 
	Data [][]byte //slice com strings em byte 
}


// RandString gera uma string aleatória em memória persistente
func RandString(n int) []byte {
	b := pmake([]byte, n) //alocação persistente 
	tx := transaction.NewRedoTx()
	tx.Begin()
	tx.Log(b)
	for i := range b {
			b[i] = byte(rand.Intn(26) + 65)
	}
	tx.End()
	transaction.Release(tx)
	return b
}

// PopulateRoot inicializa a estrutura Raiz na memória persistente
func PopulateRoot(r *Raiz) {
		tx := transaction.NewRedoTx()
		tx.Begin()
		tx.Log(r)
		r.Magic = Magic
		r.Data = pmake([][]byte,0)
		tx.End()
		transaction.Release(tx)
}

func AddString(r *Raiz) {
	s := RandString(10)

	tx := transaction.NewRedoTx()
	tx.Begin()
	tx.Log(r)
	r.Data = append(r.Data, s)
	tx.End()
	transaction.Release(tx)
}

func PrintTudo(r *Raiz) {
		for chave,valor := range r.Data{
			fmt.Printf("[%d] %s\n",chave,string(valor))
	}
}
