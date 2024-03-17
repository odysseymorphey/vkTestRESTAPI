CREATE TABLE IF NOT EXISTS actors (
	actor_id serial primary key,
	actor_name varchar(30) not null,
	gender varchar(10) not null,
	birthdate date
);

CREATE TABLE IF NOT EXISTS movies (
	movie_id serial primary key,
	title varchar(150),
	description TEXT,
	release_date date,
	rating int
);

CREATE TABLE IF NOT EXISTS relations (
	relation_id serial primary key,
	actor_id int references actors(actor_id),
	movie_id int references movies(movie_id)
);