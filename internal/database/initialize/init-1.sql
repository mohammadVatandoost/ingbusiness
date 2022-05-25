CREATE ROLE crm if not exists;
ALTER ROLE crm WITH LOGIN;
ALTER ROLE crm WITH PASSWORD 'something';
CREATE DATABASE crm_db;
GRANT ALL PRIVILEGES ON DATABASE crm_db TO assist;
ALTER USER crm_db superuser;
