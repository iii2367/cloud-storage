CREATE TABLE tree_nodes (
    id          UUID PRIMARY KEY,
    parent_id   UUID REFERENCES tree_nodes(id) ON DELETE CASCADE,
    user_id     INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    file_type   VARCHAR(16) NOT NULL CHECK (file_type IN ('file', 'folder')),
    mime_type   VARCHAR(255) NOT NULL,

    upload_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    name        VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    
    size        BIGINT NOT NULL DEFAULT 0
);
