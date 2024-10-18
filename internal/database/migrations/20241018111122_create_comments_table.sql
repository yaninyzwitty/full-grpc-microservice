-- +goose Up
-- +goose StatementBegin
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,              
    content TEXT NOT NULL,              
    post_id TEXT NOT NULL,              -- Change to TEXT to match posts.id
    user_id TEXT NOT NULL,              -- Change to TEXT to match users.id
    likes INTEGER DEFAULT 0,            
    created_at TIMESTAMP NOT NULL,  
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE, 
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments;
-- +goose StatementEnd
