CREATE TABLE users
(
    id serial not null unique,
    first_name text not null,
    last_name text not null,
    email text not null unique,
    created_at time,
    updated_at time
);

CREATE TABLE messages
(
    id serial not null UNIQUE,
    message_text TEXT NOT NULL,
    chat_id INTEGER NOT NULL,
    created_by INTEGER REFERENCES users(id) on DELETE CASCADE NOT NULL,
    created_at TIME,
    updated_at TIME
);

CREATE TABLE chats
(
    id serial NOT NULL UNIQUE,
    chat_name TEXT NOT NULL,
    chat_description TEXT,
    created_by INTEGER NOT NULL,
    created_at TIME NOT NULL,
    updated_at TIME

);

CREATE TABLE user_message
(
    user_id INTEGER REFERENCES users(id) on DELETE CASCADE not null,
    message_id INTEGER REFERENCES messages(id) on DELETE CASCADE NOT NULL,
    is_read BOOLEAN
);

CREATE TABLE user_credential
(
    id serial REFERENCES users(id) on DELETE CASCADE NOT NULL UNIQUE,
    email TEXT NOT NULL,
    password text not null
);


