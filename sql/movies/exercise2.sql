INSERT INTO  movies(title, rating, awards, release_date, length, genre_id) 
VALUES ("Resident Evil 2", 8, 2, "2004-11-19", 89, 2);

UPDATE actors SET favorite_movie_id = 22 WHERE id = 48;

CREATE TEMPORARY TABLE temp_movies
SELECT * FROM movies

DELETE FROM temp_movies 
WHERE awards < 5

SELECT a.first_name, a.last_name, m.title as favorite_movie, m.awards  
FROM actors a INNER JOIN movies m on a.favorite_movie_id = m.id 
WHERE awards > 3  

CREATE UNIQUE INDEX  title_index 
ON movies (title)

SHOW INDEXES FROM  movies 

EXPLAIN
SELECT * from movies 
WHERE title  = "intensamente"


