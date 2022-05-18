CREATE ROLE assist;
ALTER ROLE assist WITH LOGIN;
ALTER ROLE assist WITH PASSWORD 'something';
CREATE DATABASE assist_db;
GRANT ALL PRIVILEGES ON DATABASE assist_db TO assist;
ALTER USER assist superuser;
