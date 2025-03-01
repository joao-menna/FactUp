CREATE TABLE "user" (
	"id" SERIAL NOT NULL UNIQUE,
	"email" TEXT NOT NULL UNIQUE,
	"display_name" TEXT NOT NULL,
	PRIMARY KEY("id")
);

CREATE UNIQUE INDEX "user_email_index"
ON "user" ("email");

CREATE TABLE "post" (
	"id" SERIAL NOT NULL UNIQUE,
	"user_id" INTEGER NOT NULL,
	"body" TEXT NOT NULL,
	"source" TEXT,
	"image_path" TEXT,
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
ON "user_interaction" ("post_id", "user_id", "score");

ALTER TABLE "user"
ADD FOREIGN KEY("id") REFERENCES "post"("user_id")
ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "post"
ADD FOREIGN KEY("id") REFERENCES "user_interaction"("post_id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
ALTER TABLE "user"
ADD FOREIGN KEY("id") REFERENCES "user_interaction"("user_id")
ON UPDATE NO ACTION ON DELETE NO ACTION;
