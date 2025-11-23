#!/bin/bash
# Script para configurar ambiente Go no node04

NODE_NAME=$(hostname)
echo "Configurando Go no node: $NODE_NAME"

# Cria diretório no /home-ext se não existir
mkdir -p /home-ext/$USER

# Verifica se Go já está instalado localmente
if [ ! -d "/home-ext/$USER/go" ]; then
	echo "Go não encontrado. Instalando..."
	cd /home-ext/$USER
	wget https://go.dev/dl/go1.22.9.linux-amd64.tar.gz
	tar -xzf go1.22.9.linux-amd64.tar.gz
	rm go1.22.9.linux-amd64.tar.gz
	echo "Go instalado com sucesso!"
fi

# Configura variáveis de ambiente
export PATH=/home-ext/$USER/go/bin:$PATH
export GOROOT=/home-ext/$USER/go
export GOROOT_BOOTSTRAP=/home-ext/$USER/go
export GOPATH=/home-ext/$USER/go-workspace

# Adiciona ao .bashrc se ainda não existir
if ! grep -q "GOROOT_BOOTSTRAP" ~/.bashrc; then
	cat >>~/.bashrc <<'BASHRC_EOF'

# Go environment (para node04)
if [ "$(hostname)" = "node04" ]; then
    export PATH=/home-ext/$USER/go/bin:$PATH
    export GOROOT=/home-ext/$USER/go
    export GOROOT_BOOTSTRAP=/home-ext/$USER/go
    export GOPATH=/home-ext/$USER/go-workspace
fi
BASHRC_EOF
	echo "Configuração adicionada ao .bashrc"
fi

echo "Ambiente Go configurado!"
go version

echo ""
echo "Para compilar go-pmem:"
echo "  cd /home-ext/$USER"
echo "  git clone https://github.com/jerrinsg/go-pmem"
echo "  cd go-pmem"
echo "  ./make.bash"
