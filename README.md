# Finance account
Finance account é uma API REST que permite a criação de contas, autenticação e realização transferências entre elas dentro de um simples sistema bancário. A API é escrita em golang e está embasada no paradigma da clean architecture.

**Os valores monetários são tratados em centavos**

Com o uso da API é possível:
*  Criar contas
*  Listar contas criadas
*  Acessar saldo bancário pelo id da conta
*  Autenticar um usuário pelo CPF e Senha da conta
*  Realizar transferências entre usuários autenticados
*  Listar as transferências da usuária autenticada 

**Para realizar uma transferência ou lista-las o usuário precisa estar autenticado através de um token JWT**

## dependências 
* [go 1.18](https://go.dev/dl/)
* [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v4@v4.4.0)
* [docker](https://docs.docker.com/)
* [pgx](https://pkg.go.dev/github.com/jackc/pgx/v4@v4.14.1)
* [uuid](https://pkg.go.dev/github.com/google/uuid@v1.3.0)
* [mux](https://pkg.go.dev/github.com/gorilla/mux@v1.8.0)
* [migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4@v4.15.1)
* [envconfig](https://pkg.go.dev/github.com/kelseyhightower/envconfig@v1.4.0)
* [crypto](https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20210921155107-089bfa567519/bcrypt)
* [logrus](https://pkg.go.dev/github.com/sirupsen/logrus@v1.8.1)
* [pq lib](https://pkg.go.dev/github.com/lib/pq@v1.10.2)
* [dockertest](https://pkg.go.dev/github.com/ory/dockertest/v3@v3.8.1)
* [testify](https://pkg.go.dev/github.com/stretchr/testify@v1.7.1/assert)

Para baixá-las, com Go instalado na sua máquina:
```bash
go mod download
```

## Como Usar
#### Variáveis de Ambiente
Nome | Descrição | Exemplo
------------| ------------------------------------ | -------------------------------------------------
DB_URL | string com nome, usuário, porta e host do banco de dados | "postgres://postgres:4321@localhost:5432/projeto"
API_PORT | Porta em que o servidor é executado | ":5000"
TOKEN_SECRET| Segredo de geração do Token | "AjwMkrz632"
EXP_TIME | Tempo de expiração do Token | "5m" (5 minutos)

A variável de ambiente EXP_TIME é uma string de duração, pode ser definida como um número, decimal ou não e uma unidade de tempo. [Para saber mais opções de valores para essa variável](https://pkg.go.dev/time#Duration)

### Para executar via Docker-Compose:
```bash
docker-compose up --build -d
```

### Para rodar a aplicação via arquivo com comandos shell:
```bash
./run.sh
```

### Pode-se executar a aplicação através dos comandos do makefile:

Além de corrigir os códigos de importação, formata todo o código no padrão gofmt
```bash
make format
```
Expõe partes do código fora do padrão dos linters do golangci-lint
```bash
make lint
```
Executa todos os testes unitários da aplicação
```bash
make test
```
Builda o binário da aplicação
```bash
make build
```
Builda a imagem do docker da aplicação
```bash
make build-image
```
Executa localmente a aplicação
```bash
make run-local
```
## Endpoints
O corpo da resposta e da requisição estão em formato JSON

#### Accounts

**POST /accounts - criação de conta**

Corpo da requisição:

```json
{
	"name": "dan",
	"cpf": "12345678919",
	"secret": "123",
	"balance": 20
}
```

Resposta de sucesso:

* 201 Created
```json
	{
		"id": "a432498e-89b2-4827-b0e5-0bea7ac978ce",
		"name": "dan",
		"cpf": "12345678919",
		"balance": 20,
		"created_at": "2022-04-06T14:24:57Z"
	}
```

Resposta de insucesso:

* 400 Bad Request

Caso CPF informado esteja em um formato inválido
```json
{
    "reason": "cpf informed is invalid"
}
```
Caso o nome informado esteja vazio
```json
{
    "reason": "name informed is empty"
}
```

* 409 Conflict

Caso o cpf já exista no banco
```json
{
    "reason": "cpf informed alredy exists"
}
```

* 500 Internal Server Error

Caso algum erro inesperado no servidor ocorra
```json
{
    "reason": "internal server error"
}
```

**GET /accounts - listagem de contas**

Resposta de sucesso:

* 200 OK

```json
{
	"list": [
		{
			"id": "cea09752-2cee-4abb-85be-3d8b28083e32",
			"name": "dan",
			"created_at": "2022-04-14T14:07:57Z",
			"balance": 20
		}
	]
}
```

* 500 Internal Server Error

Caso algum erro inesperado no servidor ocorra
```json
{
    "reason": "internal server error"
}
```

**GET /accounts/{id}/balance - busca de saldo pelo id da conta**

* 200 OK

```json
{
	"balance": 20
}
```

* 404 Not Found

Caso a conta não tenha sido encontrada
```json
{
	"reason": "account not found"
}
```

* 500 Internal Server Error

Caso algum erro inesperado no servidor ocorra
```json
{
    "reason": "internal server error"
}
```

#### Transfers

**POST /transfers - criação da transferência**

Cabeçalho da Requisição:

{
    "Authorization" : "Bearer Token"
}

Exemplo:

```json
{
    "Authorization:" "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZWEwOTc1Mi0yY2VlLTRhYmItODViZS0zZDhiMjgwODNlMzIiLCJleHAiOjE2NDk5NTkxNTgsImlhdCI6MTY0OTk1ODg1OCwianRpIjoiZDJkY2RlMjAtOGI1ZS00ZjI0LTg2MDctZjRlMGQxNzFkZmI2In0.VQ2GZIyahFQFWIbTjtoECrbCXTgdxwTdkDu7PMH4DAM"
}
```

Corpo da requisição:

```json
{
	"account_destination_id": "c3d78e63-fad9-45e8-8248-b8cc078d1bdf",
	"amount": 10
}
```

Resposta de sucesso:

* 201 Created

```json
{
    "id": "ddc8969a-e5e3-4723-9a3e-6c4c1fc1353f",
    "account_origin_id": "cea09752-2cee-4abb-85be-3d8b28083e32",
    "account_destination_id": "030d3cdf-5e9d-4c2f-a47b-d37555c2fdd2",
    "amount": 1,
    "create_at": "2022-04-14T17:23:49Z"
}
```

Resposta de insucesso:

* 400 Bad Request

Caso a conta de origem da transferência não seja encontrada
```json
{
    "reason": "transfer origin account not found"
}
```

Caso a conta de destino da transferência não seja encontrada
```json
{
    "reason": "transfer destination account not found"
}
```

Caso a quantia seja menor ou igual a zero
```json
{
    "reason": "amount is less or equal zero"
}
```

Caso tentativa de transferência seja para a mesma conta
```json
{
    "reason": "transfer attempt to the same account"
}
```

Caso a conta de origem não tenha saldo suficiente para realizar a transferência 
```json
{
    "reason": "insufficient balance on account"
}
```

* 500 Internal Server Error

Caso algum erro inesperado no servidor ocorra
```json
{
    "reason": "internal server error"
}
```

**GET /transfers - listagem das transferências da usuária autenticada**

Cabeçalho da Requisição:

{
    "Authorization" : "Bearer Token"
}

Exemplo:

```json
{
    "Authorization:" "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZWEwOTc1Mi0yY2VlLTRhYmItODViZS0zZDhiMjgwODNlMzIiLCJleHAiOjE2NDk5NTkxNTgsImlhdCI6MTY0OTk1ODg1OCwianRpIjoiZDJkY2RlMjAtOGI1ZS00ZjI0LTg2MDctZjRlMGQxNzFkZmI2In0.VQ2GZIyahFQFWIbTjtoECrbCXTgdxwTdkDu7PMH4DAM"
}
```

Resposta de sucesso:

* 200 OK

```json
{
        "id": "ddc8969a-e5e3-4723-9a3e-6c4c1fc1353f",
        "account_origin_id": "cea09752-2cee-4abb-85be-3d8b28083e32",
        "account_destination_id": "030d3cdf-5e9d-4c2f-a47b-d37555c2fdd2",
        "amount": 1,
        "created_at": "2022-04-14 14:23:49.170126 -0300 -03"
}
```

Resposta de insucesso:

* 404 Not Found

Caso não exista transferências pertencentes aquela conta 
```json
{
    "reason": "no transfer found for this account"
}
```

* 500 Internal Server Error

Caso algum erro inesperado no servidor ocorra
```json
{
    "reason": "internal server error"
}
```

#### Login

**POST /login - autenticação do usuário**

Corpo da requisição:

```json
{
    "cpf":"12345678910",
    "secret":"123"
}
```

Resposta de sucesso:

* 201 Created

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZWEwOTc1Mi0yY2VlLTRhYmItODViZS0zZDhiMjgwODNlMzIiLCJleHAiOjE2NDk5NTkxNTgsImlhdCI6MTY0OTk1ODg1OCwianRpIjoiZDJkY2RlMjAtOGI1ZS00ZjI0LTg2MDctZjRlMGQxNzFkZmI2In0.VQ2GZIyahFQFWIbTjtoECrbCXTgdxwTdkDu7PMH4DAM"
}
```

	
Resposta de insucesso:

* 400 Bad Request

Caso um senha vazia tenha sido informada
```json
{
    "reason": "empty secret was informed"
}
```

Caso o cpf informado esteja em um formato inválido
```json
{
    "reason": "cpf informed is invalid"
}
```

* 403 Forbidden 

Caso exista alguma credencial inválida
```json
{
    "reason": "invalid credentials"
}
```

* 500 Internal Server Error

Caso algum erro inesperado no servidor ocorra
```json
{
    "reason": "internal server error"
}
```

## Licença
[MIT](https://choosealicense.com/licenses/mit/)

## Agradecimentos
Agradeço ao meu mentor [Pedro](https://github.com/pedroyremolo) e aos colegas de time por me terem me fornecido o conteúdo certo para realização da aplicação.
