docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres


## Login to postgres
sudo -u postgres psql

## See users
SELECT grantee, privilege_type, table_schema, table_name 
FROM information_schema.role_table_grants 
WHERE grantee = 'myuser';


## Create the Database
CREATE DATABASE s13f;

## Create the User (Role)
CREATE USER s13f WITH PASSWORD 's13f';

## Grant User Ownership of the Database
ALTER DATABASE s13f OWNER TO s13f;

## Grant All Privileges on the Database
GRANT ALL PRIVILEGES ON DATABASE s13f TO s13f;

## Allow User to Create Tables and Objects in the Schema
\c s13f
GRANT ALL ON SCHEMA public TO s13f;

## Verify the User’s Permissions
\du

## List ownership
SELECT datname, datdba::regrole FROM pg_database WHERE datname = 's13f';

## Change Authentication Method to Password (md5 or scram-sha-256)
By default, PostgreSQL uses peer authentication for local connections.
sudo vim /etc/postgresql/14/main/pg_hba.conf
And change the last field from peer to md5
```
local   all             all                                     peer
```
TO
```
local   all             all                                     md5
```

## Log back in as s13f
psql -U s13f -d s13f -W
