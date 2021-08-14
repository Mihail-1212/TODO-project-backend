-- name: migrate
CREATE TABLE IF NOT EXISTS task 
(
    id serial NOT NULL PRIMARY KEY,
    name VARCHAR ( 50 ) NOT NULL,
	userId INT NOT NULL REFERENCES "user"(id) ON DELETE CASCADE,
	description VARCHAR ( 255 ) NOT NULL,
	dateStart TIMESTAMP NOT NULL
)



