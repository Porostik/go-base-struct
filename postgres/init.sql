DROP TABLE IF EXISTS person;

CREATE TABLE IF NOT EXISTS person (
    id INTEGER PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL
);

INSERT INTO person (id, name, email) VALUES (
    1, 'Name1', 'email1'
);
INSERT INTO person (id, name, email) VALUES (
    2, 'Name2', 'email2'
);
INSERT INTO person (id, name, email) VALUES (
    3, 'Name3', 'email3'
);
