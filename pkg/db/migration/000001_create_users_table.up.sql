CREATE TABLE "profile" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "avatar_url" bytea,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "room" (
  "id" bigint PRIMARY KEY,
  "owner_id" bigint,
  "member_id" bigint[],
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "message" (
  "id" bigserial PRIMARY KEY,
  "room_id" bigint,
  "content" text,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "code" (
  "id" bigserial PRIMARY KEY,
  "room_id" bigint UNIQUE NOT NULL,
  "content" text,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);




CREATE INDEX ON "room" ("owner_id");

CREATE INDEX ON "message" ("room_id");

CREATE INDEX ON "code" ("room_id");

ALTER TABLE "room" ADD FOREIGN KEY ("owner_id") REFERENCES "profile" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "message" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "code" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

