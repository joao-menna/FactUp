CREATE TABLE "user" (
	"id" SERIAL NOT NULL UNIQUE,
	"email" TEXT,
	"display_name" TEXT,
	"image_path" TEXT,
	"category" TEXT NOT NULL,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"banned" BOOLEAN NOT NULL DEFAULT FALSE,
	PRIMARY KEY("id")
);


CREATE TABLE "post" (
	"id" SERIAL NOT NULL UNIQUE,
	"type" TEXT NOT NULL,
	"user_id" INTEGER NOT NULL,
	"body" TEXT NOT NULL,
	"source" TEXT,
	"image_path" TEXT,
	"created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id")
);


CREATE TABLE "user_interaction" (
	"id" SERIAL NOT NULL UNIQUE,
	"post_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	"score" SMALLINT NOT NULL,
	PRIMARY KEY("id")
);

CREATE UNIQUE INDEX "user_interaction_index_0"
ON "user_interaction" ("post_id", "user_id");

CREATE TABLE "user_bot" (
	"id" SERIAL NOT NULL UNIQUE,
	"user_id" INTEGER NOT NULL,
	"name" TEXT NOT NULL,
	"secret" TEXT NOT NULL,
	PRIMARY KEY("id")
);


ALTER TABLE "post"
ADD FOREIGN KEY("user_id") REFERENCES "user"("id")
ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "user_interaction"
ADD FOREIGN KEY("post_id") REFERENCES "post"("id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "user_interaction"
ADD FOREIGN KEY("user_id") REFERENCES "user"("id")
ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "user_bot"
ADD FOREIGN KEY("user_id") REFERENCES "user"("id")
ON UPDATE NO ACTION ON DELETE CASCADE;
