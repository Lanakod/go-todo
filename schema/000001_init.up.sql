CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE todo_lists
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users_lists
(
    id      SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    list_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists ON DELETE CASCADE
);

CREATE TABLE todo_items
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    done        BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE TABLE lists_items
(
    id      SERIAL PRIMARY KEY,
    item_id INT NOT NULL,
    list_id INT NOT NULL,
    FOREIGN KEY (item_id) REFERENCES todo_items ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists ON DELETE CASCADE
);