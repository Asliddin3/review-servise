CREATE Table if NOT exists review(
  id serial PRIMARY KEY,
  post_id int,
  customer_id int,
  description TEXT,
  review int,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);
