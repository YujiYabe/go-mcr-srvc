INSERT INTO roles (role_name, description) VALUES 
('admin', 'Administrator with full access'),
('user', 'Regular user with limited access');

INSERT INTO users (auth0_user_id, email, full_name) VALUES 
('auth0|person001', 'aaaa@gmail.com', 'a a'),
('auth0|person002', 'bbbb@gmail.com', 'b b'),
('auth0|abc123', 'user@example.com', 'John Doe'),
('auth0|xyz456', 'admin@example.com', 'Admin User');

INSERT INTO user_roles (user_id, role_id) VALUES 
(3, 2),  -- ユーザー3にロール "user" を割り当て
(4, 1);  -- ユーザー4にロール "admin" を割り当て
