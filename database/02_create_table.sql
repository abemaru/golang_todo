\c golang_todo;

CREATE TABLE todoschema.todo (
    id SERIAL PRIMARY KEY,
    title TEXT,
    status TEXT
);

INSERT INTO todoschema.todo (title, status) VALUES ('first todo', 'active');