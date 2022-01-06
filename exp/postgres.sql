CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT,
  email TEXT NOT NULL
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  user_id INT NTO NULL,
  amount INT,
  description TEXT
);