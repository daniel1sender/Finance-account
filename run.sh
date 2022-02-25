#%/bin/bash
export DB_URL="postgres://postgres:1234@localhost:5432/desafio"
export API_PORT=":3000"
export TOKEN_SECRET="123"
go run main.go
