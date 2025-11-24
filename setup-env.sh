#!/bin/bash

set -e # Encerra o script se algum comando falhar

echo "=== Configuração do Ambiente Go Local (Sem Sudo) ==="

# Diretórios importantes (todos em /home-ext/$USER)
BASE_DIR="/home-ext/$USER"
BOOTSTRAP_DIR="$BASE_DIR/go-bootstrap"
PMEM_DIR="$BASE_DIR/go-pmem"
WORKSPACE_DIR="$BASE_DIR/go-workspace"
GO_VERSION="go1.22.9.linux-amd64.tar.gz"

mkdir -p "$BASE_DIR"
mkdir -p "$WORKSPACE_DIR"

# ---------------------------------------------------------
# 1. INSTALAR GO PADRÃO (BOOTSTRAP) LOCALMENTE
# ---------------------------------------------------------
if [ ! -d "$BOOTSTRAP_DIR" ]; then
    echo "[1/4] Instalando Go Padrão (Bootstrap) em $BOOTSTRAP_DIR..."
    cd "$BASE_DIR"
    wget -q https://go.dev/dl/$GO_VERSION
    
    # Extrai o binário e renomeia para go-bootstrap
    tar -xzf $GO_VERSION
    mv go go-bootstrap
    rm $GO_VERSION
    
    echo "-> Go Bootstrap 1.22.9 instalado com sucesso."
else
    echo "[1/4] Go Bootstrap já instalado."
fi

# --- Configura o PATH para usar o Go Bootstrap nesta sessão ---
export GOROOT_BOOTSTRAP=$BOOTSTRAP_DIR
export PATH=$BOOTSTRAP_DIR/bin:$PATH
export GOPATH=$WORKSPACE_DIR

# ---------------------------------------------------------
# 2. CLONAR E PREPARAR GO-PMEM
# ---------------------------------------------------------
if [ ! -d "$PMEM_DIR" ]; then
    echo "[2/4] Clonando repositório go-pmem em $PMEM_DIR..."
    cd "$BASE_DIR"
    git clone https://github.com/jerrinsg/go-pmem.git
else
    echo "[2/4] Repositório go-pmem já existe."
fi

# ---------------------------------------------------------
# 3. COMPILAR GO-PMEM USANDO O BOOTSTRAP
# ---------------------------------------------------------
echo "[3/4] Compilando go-pmem..."
cd $PMEM_DIR/src
# O make.bash usa GOROOT_BOOTSTRAP para saber qual Go usar
./make.bash

echo "-> Compilação do Go-PMEM concluída. Binário em $PMEM_DIR/bin"

# ---------------------------------------------------------
# 4. CONFIGURAR AMBIENTE (O CORRETIVO DO GOROOT)
# ---------------------------------------------------------
echo "[4/4] Configurando aliases de ambiente no .bashrc..."

cat >>~/.bashrc <<BASHRC_EOF

# --- Go Environment Switch (Configurado para /home-ext/\$USER) ---
# Use 'usar-go-pmem' para ativar o compilador modificado.
alias usar-go-pmem='export GOROOT=$PMEM_DIR; export PATH=\$GOROOT/bin:\$GOPATH/bin:\$PATH; export GO111MODULE=off; echo "Switched to Go-PMEM (Transactional)"'

# Use 'usar-go-padrao' para voltar ao Bootstrap 1.22.9 (apenas se precisar)
alias usar-go-padrao='export GOROOT=$BOOTSTRAP_DIR; export PATH=\$GOROOT/bin:\$GOPATH/bin:\$PATH; export GO111MODULE=on; echo "Switched to Standard Go 1.22.9"'

# GOROOT_BOOTSTRAP é sempre o mesmo
export GOROOT_BOOTSTRAP=$BOOTSTRAP_DIR
export GOPATH=$WORKSPACE_DIR
# -----------------------------------------------------------------

BASHRC_EOF

echo "=== CONFIGURAÇÃO FINALIZADA COM SUCESSO! ==="

### Próximos passos após executar o script ###
echo "Execute o comando: source ~/.bashrc"
echo "Em seguida, para usar o compilador modificado, execute:"
echo "usar-go-pmem"
