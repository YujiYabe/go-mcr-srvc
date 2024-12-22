set
    client_encoding = 'UTF8';

---- create ----
CREATE TABLE persons (
    id integer PRIMARY KEY,
    name VARCHAR(20),
    mail_address VARCHAR(32),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
