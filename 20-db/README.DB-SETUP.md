# Login to postgres
```bash
sudo -i -u postgres
psql
```

## Makeshift table to hold purchases
```sql
CREATE TABLE purchase_proxy (
  id SERIAL PRIMARY KEY,
  user_id TEXT NOT NULL,
  sln TEXT NOT NULL,
  item_name TEXT NOT NULL,
  quantity FLOAT,
  units TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  bought_at TIMESTAMP,
  h3index TEXT
)
```

# Once in, create the user, database, and role
```sql
CREATE USER "shopping-list" WITH PASSWORD '';
CREATE DATABASE "shopping-list" OWNER "shopping-list";
# not needed
GRANT ALL PRIVILEGES ON DATABASE "shopping-list" TO "shopping-list";
exit
```

# Verify
```bash
psql -U "shopping-list" -d "shopping-list" -h localhost
```

# Drop tables
```sql
DROP TABLE IF EXISTS purchase;
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS sl_change;
DROP TABLE IF EXISTS shopping_list;
DROP TYPE IF EXISTS purpose_enum;
DROP TYPE IF EXISTS change_enum;
```

# Create the tables
```sql
-- 1. Create the enum type (if not already done)
CREATE TYPE purpose_enum AS ENUM ('todo', 'store', 'howto');
CREATE TYPE change_enum AS ENUM('add', 'delete', 'rename');

-- 2. Create the table
CREATE TABLE shopping_list (
  id SERIAL PRIMARY KEY,
  user_id TEXT NOT NULL,
  sln TEXT NOT NULL,
  description TEXT,
  purpose purpose_enum NOT NULL,
  h3index TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  modified_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE item (
  id SERIAL PRIMARY KEY,
  shopping_list_id INT NOT NULL REFERENCES shopping_list(id),
  item_name TEXT NOT NULL
);

CREATE TABLE sl_change (
  id SERIAL PRIMARY KEY,
  user_id TEXT NOT NULL,
  sln TEXT NOT NULL,
  change change_enum NOT NULL,
  field_name TEXT,
  old_value TEXT,
  new_value TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE purchase (
  id SERIAL PRIMARY KEY,
  shopping_list_id INT NOT NULL REFERENCES shopping_list(id),
  user_id TEXT NOT NULL,
  item_id INT NOT NULL REFERENCES item(id),
  quantity FLOAT,
  units TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  bought_at TIMESTAMP,
  h3index TEXT
);

CREATE TABLE purchase_ingest (
  id SERIAL PRIMARY KEY,
  sln TEXT NOT NULL,
  user_id TEXT NOT NULL,
  item_name TEXT NOT NULL,
  created_at TIMESTAMP,
  bought_at TIMESTAMP
);
```


## Which items were not purchased in the intended store
```sql
SELECT p.*
FROM purchase p
JOIN shopping_list s ON p.shopping_list_id = s.id
WHERE p.h3index IS DISTINCT FROM s.h3index;
```

## /api/user/slns/count
```sql
SELECT sln, count(*)
FROM purchase_ingest 
WHERE bought_at >= CURRENT_DATE - INTERVAL '50 days'
  AND bought_at < CURRENT_DATE - INTERVAL '1 day' 
GROUP BY  sln;
```

## /api/user/sln/items/count
```sql
SELECT item_name, COUNT(*)
FROM purchase_ingest
WHERE bought_at >= CURRENT_DATE - INTERVAL '300 days'
  AND bought_at < CURRENT_DATE - INTERVAL '1 day'
  AND sln = 'Moms-WholeFoods'
GROUP BY item_name;
```


## /api/user/sln/item/cadence/count
```sql
--- daily
date_trunc('day', bought_at)
date_trunc('day', bought_at) + INTERVAL '1 day' - INTERVAL '1 second'
-- weekly 
date_trunc('week', bought_at)
date_trunc('week', bought_at) + INTERVAL '6 days'
-- Monthly
date_trunc('month', bought_at)
date_trunc('month', bought_at) + INTERVAL '1 month' - INTERVAL '1 day'
-- Yearly	
date_trunc('year', bought_at)
date_trunc('year', bought_at)	+ INTERVAL '1 year' - INTERVAL '1 day'
```
```sql
SELECT 
  date_trunc('month', bought_at)::date AS month_start,
  date_trunc('month', bought_at)::date + INTERVAL '1 month' - INTERVAL '1 day' AS month_end,
  COUNT(*) AS count
FROM purchase_ingest
WHERE bought_at >= CURRENT_DATE - INTERVAL '300 days'
  AND bought_at < CURRENT_DATE - INTERVAL '1 day'
  AND item_name = 'Milk'
  AND sln = 'Moms-WholeFoods'
GROUP BY month_start
ORDER BY month_start;
```

http://localhost:6601/api/user/sln/item/cadence/count?                   t1=300&t2=0&sln=Moms-WholeFoods&item=Milk&cadence=weekly
https://www.baggybara.com/shopping-list/api/user/sln/item/cadence/count? sln=Moms-WholeFoods&item=Milk&t1=30&t2=0&cadence=weekly