-- 02-create-user.sql
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_user WHERE usename = 'postgres') THEN
        CREATE USER postgres WITH PASSWORD '4631';
        GRANT ALL PRIVILEGES ON DATABASE grdb TO postgres;
    END IF;
END $$;
