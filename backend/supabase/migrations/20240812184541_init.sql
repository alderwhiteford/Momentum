CREATE TYPE oauthProvider AS ENUM ('google');

CREATE TABLE IF NOT EXISTS users (
    user_id         uuid PRIMARY KEY,
    created_at      timestamp with time zone DEFAULT now() NOT NULL,
    provider        oauthProvider NOT NULL,
    email           text NOT NULL,
    name            text NOT NULL
);