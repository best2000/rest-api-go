--psql -U postgres
--\list

DROP DATABASE IF EXISTS app;
CREATE DATABASE app;

DROP TABLE IF EXISTS dogs;
CREATE TABLE dogs(
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 255 ) NOT NULL,
	breed VARCHAR ( 255 ) NOT NULL,
	created_at TIMESTAMP NOT NULL default current_timestamp,
	updated_at TIMESTAMP default current_timestamp
);

INSERT INTO dogs("name", breed)
VALUES ('big', 'fat bear'),
       ('brown', 'dearling'),
       ('black', 'niggas');

SELECT * FROM dogs;