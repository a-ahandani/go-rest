-- 01-create-db.sql
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'gorest') THEN
        CREATE DATABASE "gorest";
    END IF;
END $$;