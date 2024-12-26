
insert into persons values(1, 'a a', 'aaaa@gmail.com');
insert into persons values(2, 'b b', 'bbbb@gmail.com');

INSERT INTO roles (role_name, description) VALUES 
('admin', 'Administrator with full access'),
('user', 'Regular user with limited access');

INSERT INTO users (auth0_user_id, email, full_name) VALUES 
('auth0|abc123', 'user@example.com', 'John Doe'),
('auth0|xyz456', 'admin@example.com', 'Admin User');

INSERT INTO user_roles (user_id, role_id) VALUES 
(1, 2),  -- ユーザー1にロール "user" を割り当て
(2, 1);  -- ユーザー2にロール "admin" を割り当て

