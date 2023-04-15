-- CREATE ROLE crm;
DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'crm') THEN

      RAISE NOTICE 'Role "crm" already exists. Skipping.';
   ELSE
      BEGIN   -- nested block
        CREATE ROLE crm LOGIN PASSWORD 'something';
      EXCEPTION
         WHEN duplicate_object THEN
            RAISE NOTICE 'Role "crm" was just created by a concurrent transaction. Skipping.';
      END;
   END IF;
END
$do$;
--
-- ALTER ROLE crm WITH LOGIN;
-- ALTER ROLE crm WITH PASSWORD 'something';
CREATE DATABASE crm_db;
GRANT ALL PRIVILEGES ON DATABASE crm_db TO assist;
ALTER USER crm_db superuser;
