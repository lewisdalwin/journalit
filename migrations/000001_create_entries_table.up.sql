-- our initial entries table

CREATE TABLE IF NOT EXISTS entries (
    id serial PRIMARY KEY,
    content text NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);