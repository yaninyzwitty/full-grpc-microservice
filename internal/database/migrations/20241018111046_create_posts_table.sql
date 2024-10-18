-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id TEXT PRIMARY KEY,               -- Keep as TEXT or change to INTEGER
    content TEXT NOT NULL,             
    author_id TEXT NOT NULL,           -- Change to TEXT to match users.id
    likes INTEGER DEFAULT 0,           
    created_at TIMESTAMP NOT NULL,      
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
