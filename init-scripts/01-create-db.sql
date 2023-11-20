-- 01-create-db.sql
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'grdb') THEN
        CREATE DATABASE "grdb";
    END IF;
END $$;