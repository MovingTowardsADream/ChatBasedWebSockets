CREATE OR REPLACE FUNCTION make_uid() RETURNS text AS $$
DECLARE
new_uid text;
    done bool;
BEGIN
    done := false;
    WHILE NOT done LOOP
        new_uid := md5(''||now()::text||random()::text);
        done := NOT exists(SELECT 1 FROM users WHERE id=new_uid);
END LOOP;
RETURN new_uid;
END;
$$ LANGUAGE PLPGSQL VOLATILE;

CREATE TABLE IF NOT EXISTS users
(
    id TEXT DEFAULT make_uid()::text NOT NULL UNIQUE,
    email TEXT not null unique,
    username TEXT not null unique,
    password TEXT not null,
    time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS users_idx ON users(id);
