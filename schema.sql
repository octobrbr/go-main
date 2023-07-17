DROP TABLE IF EXISTS posts, authors;

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL
);

INSERT INTO authors (id, name) VALUES (0, 'Ayooluwa Isaiah');
INSERT INTO authors (id, name) VALUES (1, 'Ukeje Goodness');
INSERT INTO authors (id, name) VALUES (2, 'Soumi Bardhan');
INSERT INTO authors (id, name) VALUES (3, 'Emmanuel John');