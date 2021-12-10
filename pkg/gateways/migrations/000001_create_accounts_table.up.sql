CREATE TABLE IF NOT EXISTS accounts(
id UUID PRIMARY KEY,
name text NOT NULL,
cpf text NOT NULL,
secret text NOT NULL,
balance int NOT NULL,
createdAt TIMESTAMP WITH TIME ZONE NOT NULL
);