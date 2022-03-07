/*migration up of table tokens*/
CREATE TABLE IF NOT EXISTS tokens(
    id UUID PRIMARY KEY,
    sub TEXT NOT NULL,
    exp_time TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    token TEXT NOT NULL
);