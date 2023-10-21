# API-REDESOCIAL-GOLANG
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySql](https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)

Essa é uma API de uma rede social, onde tem os métodos de usuários, publicacões, seguidores, onde poderá seguir, parar de seguir, autenticacão, curtir e descurtir publicacões

## Conteúdo

- [Instalação](#instalação)
- [Configuração](#configuração)
- [Iniciar](#iniciar)
- [API Endpoints](#api-endpoints)
- [Autenticação](#autenticação)
- [Banco de dados](#bancodedados)

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/joaopedropc7/API-REDESOCIAL-GOLANG.git
```

2. Instale [MySQL](https://www.mysql.com/downloads/)

## Inicio

1. Inicie a aplicação com o GO RUN
2. A API ficará acessível em http://localhost:5000

## API Endpoints
Siga esses endpoints para acessar a API:

```markdown
GET /usuarios - Buscar usuarios.

GET /usuarios/{id} - Busca um usuario por id

PUT /usuarios/{usuarioId} - Atualiza um usuario pelo ID 

Delete /usuarios/{id} - Deleta um plano pelo ID 

POST /usuarios - Cadastra um usuario

POST /usuarios/{usuarioId}/seguir - Segue um usuario pelo ID

POST /usuarios/{usuarioId}/parar-de-seguir - Para de seguir um usuario pelo ID

GET /usuarios/{usuarioId}/seguidores - Busca os próprios seguidores

GET /usuarios/{usuarioId}/seguindo - Busca os usuários na qual segue pelo ID

POST /usuarios/{usuarioId}/atualizar-senha - Atualiza a senha

----------------------------------------------

POST /login - Faz login

----------------------------------------------

POST /publicacoes - Cria publicacão

GET /publicacoes - Busca publicacões

GET /publicacoes/{publicacaoId} - Busca publicacão por ID

PUT /publicacoes/{publicacaoId} - Atualiza publicacão por ID

DELETE /publicacoes/{publicacaoId} - Deleta publicacão por ID

GET /usuarios/{usuarioId}/publicacoes - Busca publicacão por usuário

POST /publicacoes/{publicacaoId}/curtir - Curte publicacão

POST /publicacoes/{publicacaoId}/descurtir - Descurtir publicacão
```

## Authentication
Esse projeto utiliza JWT TOKEN para autenticacão

## Database
Este projeto utiliza [MySQL](https://www.mysql.com/downloads/) database.


