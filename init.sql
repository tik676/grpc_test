CREATE TABLE todo(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
	description   TEXT NOT NULL,
	completed BOOL NOT NULL
);