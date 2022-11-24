CREATE TABLE users
(
    id serial not null unique,
    first_name text not null,
    last_name text not null,
    email text not null unique,
    created_at time not NULL,
    updated_at time
);

CREATE TABLE user_message
(
    user_id INTEGER not null,
    message_id INTEGER NOT NULL,
    is_read BOOLEAN
);

CREATE TABLE user_credential
(
    id serial NOT NULL UNIQUE,
    email TEXT NOT NULL,
    password text not null
);

CREATE TABLE messages
(
    id serial not null UNIQUE,
    message_text TEXT NOT NULL,
    chat_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at TIME NOT NULL,
    updated_at TIME
);

CREATE TABLE chats
(
    id serial NOT NULL UNIQUE,
    name TEXT NOT NULL,
    description TEXT,
    created_by INTEGER NOT NULL,
    created_at TIME NOT NULL,
    updated_at TIME

);
