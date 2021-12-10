CREATE TABLE IF NOT EXISTS transfers(
id UUID PRIMARY KEY,
account_origin_id UUID UNIQUE NOT NULL,
account_destination_id UUID UNIQUE NOT NULL,
amount int NOT NULL,
create_at TIMESTAMP WITH TIME ZONE NOT NULL
);