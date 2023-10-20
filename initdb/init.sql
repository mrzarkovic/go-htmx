CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL DEFAULT false
);

INSERT INTO todos (title) VALUES ('Learn SQL');
INSERT INTO todos (title) VALUES ('Learn Go');
INSERT INTO todos (title) VALUES ('Learn Docker');
INSERT INTO todos (title) VALUES ('Learn Kubernetes');
INSERT INTO todos (title) VALUES ('Learn React');
INSERT INTO todos (title) VALUES ('Learn Vue');
INSERT INTO todos (title) VALUES ('Learn Angular');
INSERT INTO todos (title) VALUES ('Learn Rust');
INSERT INTO todos (title) VALUES ('Learn Java');
INSERT INTO todos (title) VALUES ('Learn C++');
INSERT INTO todos (title) VALUES ('Learn C');