CREATE EXTENSION IF NOT EXISTS "pgcrypto";

 CREATE TABLE IF NOT EXISTS ratings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID not null,
  movie_id UUID not null,
  rating INTEGER CHECK (rating >= 1 AND rating <= 5),
  user_comment TEXT,
  created_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
)
