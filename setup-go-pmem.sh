#!/usr/bin/env bash

#Diretório onde o compilador Go modificado será instalado

INSTALL_DIR="$HOME/go-pmem"

echo "Criando diretório em $INSTALL_DIR..."
mkdir -p "$INSTALL_DIR"
cd "$INSTALL_DIR" || exit 1

# Clonando o repositório do go-pmem
echo "Clonando repositório do go-pmem..."
git clone https://github.com/vmware/go-pmem.git
cd go-pmem/src || exit 1

# Compilando o Go modificado
echo "Compilando o Go com suporte a PMEM..."
./make.bash

if [ ! -f "$INSTALL_DIR/go-pmem/bin/go" ]; then
  echo "X Erro: compilação falhou."
  exit 1
fi

# Adicionando o alias no shell
ALIAS_CMD="alias go-pmem='$INSTALL_DIR/go-pmem/bin/go'"

# Detectar shell
if [[ "$SHELL" == */zsh ]]; then
  SHELL_RC="$HOME/.zshrc"
else
  SHELL_RC="$HOME/.bashrc"
fi

# Evita Duplicatas
if ! grep -Fxq "$ALIAS_CMD" "$SHELL_RC";then 
  echo "Adicionando alias ao $SHELL_RC"
  echo "$ALIAS_CMD" >> "$SHELL_RC"
else 
  "Alias já existe em $SHELL_RC"
fi 

echo ""
echo "go-pmem instalado com sucesso!"
echo "-> Abra um novo terminal ou rode: "
echo "    source $SHELL_RC"
echo "Agora você pode usar 'go-pmem' para compilar com suporte a memória persistente."
