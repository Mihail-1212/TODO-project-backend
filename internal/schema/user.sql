-- name: migrate

CREATE TABLE IF NOT EXISTS "user"
(
    id serial NOT NULL PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
	"password" VARCHAR ( 255 ) NOT NULL
);
