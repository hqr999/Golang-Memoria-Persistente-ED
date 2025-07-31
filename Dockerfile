FROM ubuntu:20.04 

ENV DEBIAN_FRONTEND=noninteractive
ENV LANG=C.UTF-8

#1. Atualizações e dependências do sistema 
RUN apt-get update && apt-get install -y \
  build-essential \ 
  curl \
  wget \
  git \
  vim \
  htop \
  gdb \
  make \
  autoconf \
  automake \
  libtool \
  libgmp-dev \
  libncurses-dev \
  libtinfo-dev \
  libnuma-dev \
  libatomic1 \
  zlib1g-dev \
  libpmem-dev \
  libpmemobj-dev \
  libpmemobj++-dev \
  python3 \
  python3-pip \
  diffutils \
  ca-certificates \
  xz-utils \
  libffi-dev \
  && rm -rf /var/lib/apt/lists/*


#2. Instalação do GHC 8.6.5 
WORKDIR /root 
RUN wget https://downloads.haskell.org/~ghc/8.6.5/ghc-8.6.5-x86_64-fedora27-linux.tar.xz && \
  tar -xf ghc-8.6.5-x86_64-fedora27-linux.tar.xz && \
  cd ghc-8.6.5 && \
  ./configure && \
  make install && \
  cd .. && \
  rm -rf ghc-8.6.5 ghc-8.6.5-x86_64-fedora27-linux.tar.xz

# 3. Instalação do Cabal 3.0.0
RUN wget https://downloads.haskell.org/~cabal/cabal-install-3.0.0.0/cabal-install-3.0.0.0-x86_64-unknown-linux.tar.xz && \
    tar -xf cabal-install-3.0.0.0-x86_64-unknown-linux.tar.xz && \
    mv cabal /usr/local/bin/ && \
    chmod +x /usr/local/bin/cabal && \
    rm cabal-install-3.0.0.0-x86_64-unknown-linux.tar.xz

# 4. Instalação do Stack (versão 1.9.3 via upgrade)
RUN curl -sSL https://get.haskellstack.org/ | sh && \
    stack upgrade --binary-version 1.9.3


# 5. Instalação de Happy 1.19.12 e Alex 3.2.5 via cabal
RUN cabal update && \
    cabal install happy-1.19.12 alex-3.2.5 \
      --installdir=/usr/local/bin \
      --install-method=copy \
      --overwrite-policy=always

# 6. Instalação do Neovim (versão 0.10.0 para compatibilidade com AstroNvim)
RUN curl -LO https://github.com/neovim/neovim/releases/download/v0.10.0/nvim-linux64.tar.gz && \
    tar -xzf nvim-linux64.tar.gz && \
    mv nvim-linux64 /opt/nvim && \
    ln -s /opt/nvim/bin/nvim /usr/bin/nvim && \
    rm nvim-linux64.tar.gz

# 7. Instalação do LazyVim
RUN git clone https://github.com/LazyVim/starter /root/.config/nvim && \
    rm -rf /root/.config/nvim/.git


# 8. Configuração do LazyVim para Haskell
RUN mkdir -p /root/.config/nvim/lua/plugins && \
    cat > /root/.config/nvim/lua/plugins/haskell.lua << 'EOF'
return {
  -- Haskell syntax highlighting
  {
    "neovimhaskell/haskell-vim",
    ft = "haskell",
  },
  -- Simple Haskell LSP support
  {
    "neovim/nvim-lspconfig",
    opts = {
      servers = {
        hls = {
          filetypes = { "haskell", "lhaskell" },
        },
      },
    },
  },
}
EOF

# 9. Instalação do GHCup para melhor compatibilidade com HLS
RUN curl --proto '=https' --tlsv1.2 -sSf https://get-ghcup.haskell.org | sh && \
    echo 'source /root/.ghcup/env' >> /root/.bashrc && \
    /root/.ghcup/bin/ghcup install hls --set


# 10. Configuração do ambiente
ENV TERM=xterm-256color
ENV EDITOR=nvim


# 11. Verificação das versões
RUN stack --version && \
    cabal --version && \
    ghc --version && \
    happy --version && \
    alex --version && \
    nvim --version
