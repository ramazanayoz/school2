

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
    from directors join movie_directors on directors.director_id = movie_directors.director_id
	where director_name = "Quentin Tarantino"
);


# 6. Show the thriller films whose budget is greater than 25 million$.	 
select title
from movies join genres on movies.movie_id = genres.movie_id
where budget > 25000000 and genre_name = "Thriller";



# 7. Show the drama films whose language is Italian and produced between 1990-2000.	
 select title
 from movies 
 where year between 1990 and 2000 and movie_id in(   #>1990 or year <2000 
 select genres.movie_id   #niye sadece movie_id yazmadık çünkü yukardaki tabloyu incele
 from genres join languages on genres.movie_id= languages.movie_id
 where genre_name = "Drama" and language_name = "Italian" );
 
 
 
# 8. Show the films that Tom Hanks has act and have won more than 3 Oscars.	 
select title
from movies
where oscars > 3 and movie_id in (
select movie_id
from stars join movie_stars on ( stars.star_id= movie_stars.star_id)
where star_name = "Tom Hanks"
);

# 9. Show the history films produced in USA and whose duration is between 100-200 minutes.
select title
from movies join genres on movies.movie_id = genres.movie_id
where duration between 100 and 200 and genre_name = "History" and movies.movie_id in (
	select movie_id
    from countries join producer_countries on countries.country_id = producer_countries.country_id
    where country_name = "USA"
);



# 10.Compute the average budget of the films directed by Peter Jackson.
select avg(budget)
from movies
where movie_id in (
 select movie_id
 from directors join movie_directors on directors.director_id = movie_directors.director_id
 where director_name = "Peter Jackson"
 );



# 11.Show the Francis Ford Coppola film that has the minimum budget.
select title, budget      #kısaa yol min(budget)
from movies
where movie_id in (
 select movie_id
 from directors join movie_directors on directors.director_id = movie_directors.director_id
 where director_name = "Francis Ford Coppola"   )
 order by budget
 limit 1
 ;




# 12.Show the film that has the most vote and has been produced in USA.
select title, min(votes)
 from movies
 where movie_id in (
	select movie_id
	from countries join producer_countries on countries.country_id = producer_countries.country_id
    where country_name = "USA"
 
 )
 order by votes desc  ;



#genresi direk bağladık contriesi niye direk bağlaamadık moviese
