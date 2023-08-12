
CREATE TABLE "accounts" (
  "id" serial PRIMARY KEY,
  "username" text UNIQUE NOT NULL,
  "password" text NOT NULL,
  "role" text DEFAULT 'user'
);

INSERT INTO "accounts" ("username", "password", "role") VALUES ('root', 'password', 'admin');
INSERT INTO "accounts" ("username", "password", "role") VALUES ('guest', 'password', 'user');
