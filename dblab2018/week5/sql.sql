

# 1. Show the films whose budget is greater than 10 million$ and ranking is less than 6.
select title
from movies
where budget > 1000000 and ranking < 6;


# 2. Show the action films whose rating is greater than 8.8 and produced after 2009.
select title
from movies
where rating > 8.8 and year < 2009 and movie_id in (
	select movie_id
    from genres 
    where genre_name = "action"
);



# 3. Show the drama films whose duration is more than 150 minutes and oscars is more than 2.
select title, duration, oscars
from movies 
where duration > 150 and oscars > 2 and movie_id in (
select movie_id
from genres
where genre_name = "Drama"
);

# 4. Show the films that Orlando Bloom and Ian McKellen have act together and has more than 2 Oscars.
select title
from movies
where oscars > 2 and movie_id in(
	select movie_id
    from movie_stars join stars on movie_stars.star_id = stars.star_id
    where star_name = "Orlando Bloom" and movie_id in(
		select movie_id
    from movie_stars join stars on movie_stars.star_id = stars.star_id
    where star_name = "Ian McKellen"
    
    )
    
    );



# 5. Show the Quentin Tarantino films which have more than 500000 votes and produced before 2000.	 
select title
from movies
where votes > 500000 and year < 2000 and movie_id in( 
	select movie_id
    from movie_directors join directors on movie_directors.director_id = directors.director_id
	where director_name = "Quentin Tarantino"
);


# 6. Show the thriller films whose budget is greater than 25 million$.	 

# 7. Show the drama films whose language is Italian and produced between 1990-2000.	
 
# 8. Show the films that Tom Hanks has act and have won more than 3 Oscars.	 

# 9. Show the history films produced in USA and whose duration is between 100-200 minutes.

# 10.Compute the average budget of the films directed by Peter Jackson.

# 11.Show the Francis Ford Coppola film that has the minimum budget.

# 12.Show the film that has the most vote and has been produced in USA.