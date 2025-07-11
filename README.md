# Como Rodar 

Devemos rodar os seguintes comandos para termos o conteiner:

* Rode o script bash com: **./get-pm.sh (antes faça chmod +x get-pm.sh)**

* Crie a imagem com **docker compose build**

* Rode o conteiner com **docker compose run ghc-pstm /bin/bash**

* Dentro do conteiner entre na pasta do ghc_pm e faça: 
    
    * cd /root/ghc_pm
    
    * git init
    
    * git remote add origin https://gitlab.rlp.net/zdvresearch/ghc_pm.git

