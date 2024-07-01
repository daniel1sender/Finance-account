# Finance account

Finance account is a REST API that allows the creation of accounts, authentication, and money transfers between them within a simple banking system. This API is written in Golang and is based on the clean architecture paradigm.

**Monetary values are handled in cents.**

With the use of the API, it is possible to:
*  Create accounts
*  List created accounts
*  Access account balance by account ID
*  Authenticate a user by their CPF and account password
*  Perform transfers between authenticated users
*  List transfers made by the authenticated user 

**To perform a transfer or list them, the user needs to be authenticated using a JWT token.**

## Dependencies

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
* [dockertest](https://pkg.go.dev/github.com/ory/dockertest/v3@v3.8.1)
* [testify](https://pkg.go.dev/github.com/stretchr/testify@v1.7.1/assert)

To download them, with Go installed on your machine:
```bash
go mod download
```

## How to Use

#### Environment Variables

Name | Description | Example
------------| ------------------------------------ | -------------------------------------------------
DB_URL | String with database name, user, port, and host | "postgres://postgres:4321@localhost:5432/projeto"
API_PORT | Port where the server runs | ":5000"
TOKEN_SECRET| Token generation secret | "AjwMkrz632"
EXP_TIME | Token expiration time | "5m" (5 minutes)

The EXP_TIME environment variable is a duration string and can be defined as a number, decimal or not, followed by a unit of time. [For more options for this variable, see here](https://pkg.go.dev/time#Duration)

### To run via Docker-Compose:

```bash
docker-compose up --build -d
```

### To run the application via a shell script:

```bash
./run.sh
```

### You can also run the application using commands from the Makefile:

Ensure all code is formatted according to gofmt standards
```bash
make format
```
Identify and fix code issues flagged by golangci-lint linters
```bash
make lint
```
Execute all unit tests of the application
```bash
make test
```
Build the application binary
```bash
make build
```
Build the Docker image of the application
```bash
make build-image
```
Run the application locally
```bash
make run-local
```
## Endpoints

Request and response bodies are in JSON format.

#### Accounts

**POST /accounts - account creation**

Request Body:

```json
{
	"name": "dan",
	"cpf": "12345678919",
	"secret": "123",
	"balance": 20
}
```

Successful Response:

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

Unsuccessful Response:

* 400 Bad Request

If CPF format is invalid:
```json
{
    "reason": "cpf informed is invalid"
}
```
If name is empty:
```json
{
    "reason": "name informed is empty"
}
```

* 409 Conflict

If CPF already exists in the database:
```json
{
    "reason": "cpf informed alredy exists"
}
```

* 500 Internal Server Error

If an unexpected server error occurs:
```json
{
    "reason": "internal server error"
}
```

**GET /accounts - list accounts**

Successful Response:

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

If an unexpected server error occurs:
```json
{
    "reason": "internal server error"
}
```

**GET /accounts/{id}/balance - get balance by account ID**

* 200 OK

```json
{
	"balance": 20
}
```

* 404 Not Found

If the account is not found:
```json
{
	"reason": "account not found"
}
```

* 500 Internal Server Error

If an unexpected server error occurs:
```json
{
    "reason": "internal server error"
}
```

#### Transfers

**POST /transfers - create transfer**

Request Header::

{
    "Authorization" : "Bearer Token"
}

Example:

```json
{
    "Authorization:" "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZWEwOTc1Mi0yY2VlLTRhYmItODViZS0zZDhiMjgwODNlMzIiLCJleHAiOjE2NDk5NTkxNTgsImlhdCI6MTY0OTk1ODg1OCwianRpIjoiZDJkY2RlMjAtOGI1ZS00ZjI0LTg2MDctZjRlMGQxNzFkZmI2In0.VQ2GZIyahFQFWIbTjtoECrbCXTgdxwTdkDu7PMH4DAM"
}
```

Request Body:

```json
{
	"account_destination_id": "c3d78e63-fad9-45e8-8248-b8cc078d1bdf",
	"amount": 10
}
```

Successful Response:

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

Unsuccessful Response:

* 400 Bad Request

If the origin account is not found:
```json
{
    "reason": "transfer origin account not found"
}
```

If the destination account is not found:
```json
{
    "reason": "transfer destination account not found"
}
```

If the amount is less than or equal to zero:
```json
{
    "reason": "amount is less or equal zero"
}
```

If the transfer attempt is to the same account:
```json
{
    "reason": "transfer attempt to the same account"
}
```

If the origin account has insufficient balance: 
```json
{
    "reason": "insufficient balance on account"
}
```

* 500 Internal Server Error

If an unexpected server error occurs:
```json
{
    "reason": "internal server error"
}
```

**GET /transfers - list transfers of authenticated user**

Request Header:

{
    "Authorization" : "Bearer Token"
}

Example:

```json
{
    "Authorization:" "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZWEwOTc1Mi0yY2VlLTRhYmItODViZS0zZDhiMjgwODNlMzIiLCJleHAiOjE2NDk5NTkxNTgsImlhdCI6MTY0OTk1ODg1OCwianRpIjoiZDJkY2RlMjAtOGI1ZS00ZjI0LTg2MDctZjRlMGQxNzFkZmI2In0.VQ2GZIyahFQFWIbTjtoECrbCXTgdxwTdkDu7PMH4DAM"
}
```

Successful Response:

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

Unsuccessful Response:

* 404 Not Found

If no transfers are found for the account:
```json
{
    "reason": "no transfer found for this account"
}
```

* 500 Internal Server Error

If an unexpected server error occurs:
```json
{
    "reason": "internal server error"
}
```

#### Login

**POST /login - user authentication**

Request Body::

```json
{
    "cpf":"12345678910",
    "secret":"123"
}
```

Successful Response:

* 201 Created

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZWEwOTc1Mi0yY2VlLTRhYmItODViZS0zZDhiMjgwODNlMzIiLCJleHAiOjE2NDk5NTkxNTgsImlhdCI6MTY0OTk1ODg1OCwianRpIjoiZDJkY2RlMjAtOGI1ZS00ZjI0LTg2MDctZjRlMGQxNzFkZmI2In0.VQ2GZIyahFQFWIbTjtoECrbCXTgdxwTdkDu7PMH4DAM"
}
```

	
Unsuccessful Response:

* 400 Bad Request

If an empty secret is provided:
```json
{
    "reason": "empty secret was informed"
}
```

If the CPF format is invalid:
```json
{
    "reason": "cpf informed is invalid"
}
```

* 403 Forbidden 

If any credentials are invalid:
```json
{
    "reason": "invalid credentials"
}
```

* 500 Internal Server Error

If an unexpected server error occurs:
```json
{
    "reason": "internal server error"
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Acknowledgements

Thanks to my mentor [Pedro](https://github.com/pedroyremolo) and my teammates for providing the right content to develop this application.
