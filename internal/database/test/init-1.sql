CREATE ROLE testuser;
ALTER ROLE testuser WITH LOGIN;
ALTER ROLE testuser WITH PASSWORD 'testuser';
CREATE DATABASE test_db;
GRANT ALL PRIVILEGES ON DATABASE test_db TO testuser;
ALTER USER testuser superuser;
