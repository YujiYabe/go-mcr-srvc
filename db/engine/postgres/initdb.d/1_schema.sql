set client_encoding = 'UTF8';

CREATE TABLE users (
    id SERIAL PRIMARY KEY,                          -- 自動増加の一意なID
    auth0_user_id VARCHAR(255) UNIQUE,              -- Auth0のユーザーID (例: "auth0|123456789")
    email VARCHAR(255) UNIQUE NOT NULL,             -- ユーザーのメールアドレス
    full_name VARCHAR(255),                         -- ユーザーのフルネーム
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 更新日時
);

CREATE TABLE companies (
    id SERIAL PRIMARY KEY,                          -- 会社ID
    company_code VARCHAR(50) UNIQUE NOT NULL,       -- 会社コード
    name VARCHAR(255) NOT NULL,                     -- 表示名
    legal_name VARCHAR(255),                        -- 登記上の正式名称
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 更新日時
);

CREATE TABLE office_locations (
    id SERIAL PRIMARY KEY,                                  -- 拠点ID
    company_id INT NOT NULL REFERENCES companies(id),       -- 会社ID
    location_code VARCHAR(50) NOT NULL,                     -- 拠点コード
    name VARCHAR(255) NOT NULL,                             -- 拠点名
    address TEXT,                                           -- 所在地
    timezone VARCHAR(64) DEFAULT 'Asia/Tokyo',              -- タイムゾーン
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- 更新日時
    UNIQUE (company_id, location_code)
);

CREATE TABLE departments (
    id SERIAL PRIMARY KEY,                                      -- 部署ID
    company_id INT NOT NULL REFERENCES companies(id),           -- 会社ID
    parent_department_id INT REFERENCES departments(id),         -- 親部署ID
    department_code VARCHAR(50) NOT NULL,                       -- 部署コード
    name VARCHAR(255) NOT NULL,                                 -- 部署名
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,             -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,             -- 更新日時
    UNIQUE (company_id, department_code)
);

CREATE TABLE positions (
    id SERIAL PRIMARY KEY,                                  -- 役職ID
    company_id INT NOT NULL REFERENCES companies(id),       -- 会社ID
    position_code VARCHAR(50) NOT NULL,                    -- 役職コード
    title VARCHAR(255) NOT NULL,                           -- 役職名
    rank_level INT,                                        -- 職位レベル
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- 更新日時
    UNIQUE (company_id, position_code)
);

CREATE TABLE user_employments (
    id SERIAL PRIMARY KEY,                                      -- 所属ID
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- ユーザーID
    company_id INT NOT NULL REFERENCES companies(id),           -- 会社ID
    department_id INT REFERENCES departments(id),                -- 部署ID
    position_id INT REFERENCES positions(id),                    -- 役職ID
    office_location_id INT REFERENCES office_locations(id),      -- 拠点ID
    employee_code VARCHAR(50),                                  -- 社員番号
    employment_type VARCHAR(50) DEFAULT 'full_time',             -- 雇用区分
    joined_on DATE,                                             -- 入社日
    left_on DATE,                                               -- 退職日
    is_primary BOOLEAN DEFAULT true,                            -- 主所属かどうか
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,              -- 作成日時
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,              -- 更新日時
    UNIQUE (company_id, employee_code)
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
