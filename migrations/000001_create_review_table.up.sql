CREATE Table if NOT exists review(
  id serial PRIMARY KEY,
  post_id int,
  customer_id int,
  description TEXT,
  review int
);
