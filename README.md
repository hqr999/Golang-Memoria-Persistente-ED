# Instalando o Compilador

Para ter acesso ao compilador go com memória persistente no seu sistema, rode setup-go-pmem.sh.

Para tornar o script executável rode:

* sudo chmod +x setup-go-pmem.sh 
* ./setup-go-pmem.sh 

# Rodando a lista ligada 

Faça os seguintes passos para rodar as estruturas de dados:

*   Primeiro entramos na pasta com a estrutra de dados: **cd estrutura-dados/Lista_Ligada**
*   Depois compilamos nosso código com esse comando que usa o compilador go com o heap-persistente:**GO11MODULE=off ~/go-pmem/bin/go build -txn lista_ligada_pmem.go** 

*  Executamos o arquivo passando a flag com o nome do arquivo que será mapeado na PMEM: **/lista_ligada_pmem -file=lista-ligada.goPool**

* Monte uma lista ligada simples

* Depois dê o comando para sair do programa, e então execute novamnete e voilá!! você verá que o conteúdo da lista persistente ainda estará lá!!

* Se quiser resetar sua lista ligada, basta fazer o comando: **rm lista-ligada.goPool**. Esse comando remove o arquivo com PMEM mapeada e assim você reseta sua lista.

## Observação 
Seu sistema deve estar em um linux com procesador amd64 para o compilador rodar.

Cria um alias para o comando que cria um executável para não ter tanto trabalho

