CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE NOT NULL,
  
  name TEXT NOT NULL,

  email         TEXT,
  password_hash BYTEA NOT NULL,
  salt          BYTEA NOT NULL,

  username_base TEXT,
  username_id   INT,
  CONSTRAINT username_unique UNIQUE (username_base, username_id),

  created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
