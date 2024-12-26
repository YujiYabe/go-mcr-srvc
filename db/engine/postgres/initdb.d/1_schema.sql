set client_encoding = 'UTF8';

---- create ----
CREATE TABLE persons (
    id integer PRIMARY KEY,
    name VARCHAR(20),
    mail_address VARCHAR(32),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,                          -- 自動増加の一意なID
    auth0_user_id VARCHAR(255) UNIQUE NOT NULL,     -- Auth0のユーザーID (例: "auth0|123456789")
    email VARCHAR(255) UNIQUE NOT NULL,             -- ユーザーのメールアドレス
    full_name VARCHAR(255),                         -- ユーザーのフルネーム
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 更新日時
);


CREATE TABLE roles (
    id SERIAL PRIMARY KEY,                 -- 自動増加の一意なID
    role_name VARCHAR(50) UNIQUE NOT NULL, -- ロール名 (例: "admin", "user")
    description TEXT                       -- ロールの説明
);


CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,                              -- 自動増加の一意なID
    user_id INT REFERENCES users(id) ON DELETE CASCADE, -- ユーザーID
    role_id INT REFERENCES roles(id) ON DELETE CASCADE  -- ロールID
);

