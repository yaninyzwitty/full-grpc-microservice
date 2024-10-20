-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id TEXT PRIMARY KEY,              
    username VARCHAR(50) NOT NULL UNIQUE, 
    name VARCHAR(100) NOT NULL,         
    email VARCHAR(100) NOT NULL UNIQUE, 
    bio TEXT,                           
    image_url VARCHAR(255),             
    created_at TIMESTAMP NOT NULL 
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd


