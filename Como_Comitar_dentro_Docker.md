# Comitando dentro de um conteiner

## 1. Criar um token de Acesso no Github(use tokens classic)
    Lembre-se de guardar o seu token!!!

    leia: https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-user-access-token-for-a-github-app
## 2. Configurar o Git no Container 
    git config --global user.name "Nome"   

    git config --global user.email "Seu email"

    git config --global credential.helper store

##  3. Clone o Repositório Privado
    git clone https://SEU_TOKEN@github.com/usuario/repositorio-privado.git

## 4. Agora você pode fazer comandos git nesse repositório normalmente
