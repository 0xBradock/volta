CREATE TABLE "entry" (
  "id" uuid PRIMARY KEY,
  "domain" varchar(255) UNIQUE NOT NULL,
  "login" varchar(255),
  "password" varchar(255) NOT NULL,
  "created" timestamptz NOT NULL DEFAULT (now()),
  "updated" timestamptz NOT NULL DEFAULT (now()),
  "meta" varchar
);

CREATE INDEX ON "entry" ("domain");

COMMENT ON COLUMN "entry"."domain" IS 'URL domain';

COMMENT ON COLUMN "entry"."login" IS 'login/user field content';

COMMENT ON COLUMN "entry"."password" IS 'hashed password';

COMMENT ON COLUMN "entry"."created" IS 'date of creation';

COMMENT ON COLUMN "entry"."updated" IS 'date of last update';
