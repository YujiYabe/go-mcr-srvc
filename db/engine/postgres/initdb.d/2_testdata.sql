INSERT INTO roles (role_name, description) VALUES 
('admin', 'Administrator with full access'),
('user', 'Regular user with limited access');

INSERT INTO users (auth0_user_id, email, full_name) VALUES 
('auth0|user001', 'aaaa@gmail.com', 'a a'),
('auth0|user002', 'bbbb@gmail.com', 'b b'),
('auth0|abc123', 'user@example.com', 'John Doe'),
('auth0|xyz456', 'admin@example.com', 'Admin User');

INSERT INTO companies (company_code, name, legal_name) VALUES
('acme', 'Acme Japan', 'Acme Japan Inc.');

INSERT INTO office_locations (company_id, location_code, name, address, timezone) VALUES
(1, 'tokyo-hq', 'Tokyo Headquarters', 'Tokyo, Japan', 'Asia/Tokyo'),
(1, 'osaka-branch', 'Osaka Branch', 'Osaka, Japan', 'Asia/Tokyo');

INSERT INTO departments (company_id, parent_department_id, department_code, name) VALUES
(1, NULL, 'corp', 'Corporate'),
(1, 1, 'engineering', 'Engineering'),
(1, 1, 'sales', 'Sales'),
(1, 2, 'platform', 'Platform');

INSERT INTO positions (company_id, position_code, title, rank_level) VALUES
(1, 'member', 'Member', 1),
(1, 'manager', 'Manager', 3),
(1, 'director', 'Director', 5);

INSERT INTO user_employments (
    user_id,
    company_id,
    department_id,
    position_id,
    office_location_id,
    employee_code,
    employment_type,
    joined_on,
    is_primary
) VALUES
(1, 1, 2, 1, 1, 'E001', 'full_time', '2024-04-01', true),
(2, 1, 3, 1, 2, 'E002', 'full_time', '2024-04-01', true),
(3, 1, 4, 2, 1, 'E003', 'contractor', '2024-06-01', true),
(4, 1, 1, 3, 1, 'E004', 'full_time', '2023-10-01', true);

INSERT INTO user_roles (user_id, role_id) VALUES 
(3, 2),  -- ユーザー3にロール "user" を割り当て
(4, 1);  -- ユーザー4にロール "admin" を割り当て

INSERT INTO validation_word_rules (target_type, is_blacklist, word, match_type, enabled) VALUES
('name', true, 'root', 'contains', true),
('name', true, '禁止語', 'contains', true),
('name', false, 'Admin User', 'exact', true);
