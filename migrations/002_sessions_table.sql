CREATE TABLE sessions (
    id              UUID PRIMARY KEY,
    user_id         INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash      TEXT NOT NULL UNIQUE,
    expires_at      TIMESTAMPTZ NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
    last_used_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
    revoked         BOOLEAN NOT NULL DEFAULT FALSE,
    user_agent      TEXT,
    ip_address      INET NOT NULL
)
