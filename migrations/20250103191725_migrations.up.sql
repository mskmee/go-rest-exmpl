CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  username varchar(255) NOT NULL,
  password_hash varchar(255) NOT NULL
);

CREATE TABLE todo_items (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title varchar(255) NOT NULL,
  description varchar(255),
  done boolean NOT NULL DEFAULT false
);

CREATE TABLE todo_lists (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title varchar(255) NOT NULL,
  description varchar(255) NOT NULL
);

CREATE TABLE users_lists (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  list_id UUID NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);

CREATE TABLE list_items (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  list_id UUID NOT NULL,
  item_id UUID NOT NULL,
  FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE,
  FOREIGN KEY (item_id) REFERENCES todo_items(id) ON DELETE CASCADE
);
