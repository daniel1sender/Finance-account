CREATE TABLE IF NOT EXISTS transfers(
id UUID PRIMARY KEY,
account_origin_id TEXT NOT NULL,
account_destination_id TEXT NOT NULL,
amount BIGINT NOT NULL,
created_at TIMESTAMP WITH TIME ZONE NOT NULL
);