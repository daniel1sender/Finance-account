CREATE TABLE IF NOT EXISTS transfers(
id UUID PRIMARY KEY,
account_origin_id UUID NOT NULL,
account_destination_id UUID NOT NULL,
amount BIGINT NOT NULL,
created_at TIMESTAMP WITH TIME ZONE NOT NULL
);