CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "accounts" (
  "id" uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  "username" text,
  "password" text,
  "role" text
);

INSERT INTO "accounts" ("username", "password", "role") VALUES ('root', 'password', 'admin');
INSERT INTO "accounts" ("username", "password", "role") VALUES ('guest', 'password', 'user');
