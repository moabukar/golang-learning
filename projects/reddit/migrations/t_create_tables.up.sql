CREATE TABLE threads (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL, 
);

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    thread_id UUID NOT NULL REFERENCES thread (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    votes INT DEFAULT 0
);

CREATE TABLE comments(
    id UUID PRIMARY KEY,
    post_id UUID NOT NULL REFERENCES post (id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    votes INT DEFAULT 0
);