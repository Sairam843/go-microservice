CREATE EXTENSION IF NOT EXISTS "pgcrypto";

 CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  pref_name varchar,
  email varchar UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  role TEXT CHECK (role IN ('viewer', 'admin', 'superadmin')),
  created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

