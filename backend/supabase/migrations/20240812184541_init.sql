CREATE TABLE IF NOT EXISTS users (
    user_id uuid PRIMARY KEY,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL
);