-- Creates the user
DO
$do$
BEGIN
    IF NOT EXISTS (SELECT * FROM pg_roles WHERE rolname = 'chater') THEN
        CREATE USER chater WITH PASSWORD 'password';
    END IF;
END;
$do$

-- Creates the database
CREATE database gtexter;
GRANT ALL PRIVILEGES ON DATABASE gtexter to gtexter;