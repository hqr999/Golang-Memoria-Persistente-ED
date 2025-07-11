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

# 6. Verificação das versões
RUN stack --version && \
    cabal --version && \
    ghc --version && \
    happy --version && \
    alex --version
