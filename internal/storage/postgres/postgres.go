package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/dto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Storage struct {
	log    *zap.SugaredLogger
	cancel context.CancelFunc
	db     *pgxpool.Pool
}

func NewStorage(log *zap.SugaredLogger, dsn string) (*Storage, error) {
	if log == nil {
		return nil, errors.WithMessage(errors.New("Invalid parameters"), "Empty logger")
	}

	if dsn == "" {
		return nil, errors.WithMessage(errors.New("Invalid parameters"), "Empty dsn")
	}

	db := &Storage{
		log: log,
	}

	ctx, cancel := context.WithCancel(context.Background())
	db.cancel = cancel

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection pinged")

	db.db = conn

	return db, nil
}

func (s *Storage) Close() {
	s.db.Close()
	s.cancel()
}

func (s *Storage) CreateActor(ctx context.Context, actor dto.Actor) error {
	query := `INSERT INTO actors (actor_name, gender, birthdate)
			  SELECT $1, $2, $3
			  WHERE NOT EXISTS(
			      SELECT * FROM actors WHERE actor_name = $1
			                           AND gender = $2
			                           AND birthdate = $3
			  )`

	_, err := s.db.Exec(ctx, query, actor.Name, actor.Gender, actor.BirthDate)
	if err != nil {
		s.log.Error(nil)
	}

	return err
}

func (s *Storage) UpdateActor(ctx context.Context, actor dto.Actor) error {
	query := `UPDATE actors
			  SET actor_name = $2, gender = $3, birthdate = $4
			  WHERE actor_id = $1`

	_, err := s.db.Exec(ctx, query, actor.ID, actor.Name, actor.Gender, actor.BirthDate)
	if err != nil {
		s.log.Error(nil)
	}

	return err
}

func (s *Storage) DeleteActor(ctx context.Context, id int) error {
	query := `DELETE FROM relations
			  WHERE actor_id = $1`

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		s.log.Error(err)
	}

	query = `DELETE FROM actors
			 WHERE actor_id = $1`

	_, err = s.db.Exec(ctx, query, id)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *Storage) CreateMovie(ctx context.Context, movie dto.Movie) error {
	insertActor := `INSERT INTO actors (actor_name, gender, birthdate)
					 SELECT $1, $2, $3
					 WHERE NOT EXISTS(
					 	SELECT * FROM actors WHERE actor_name = $1::varchar(30)
					 	                       AND gender = $2::varchar(10)
					 	                       AND birthdate = $3::date
					 )
					 RETURNING actor_id`

	var actorsID []int
	for _, v := range movie.Actors {
		var id int
		err := s.db.QueryRow(ctx, insertActor, v.Name, v.Gender, v.BirthDate).Scan(&id)
		if err != nil {
			s.log.Error(err)
		}

		actorsID = append(actorsID, id)
	}

	insertMovie := `INSERT INTO movies (title, description, release_date, rating)
					SELECT $1, $2, $3, $4
					WHERE NOT EXISTS(
					    SELECT * FROM movies WHERE title = $1::varchar(150)
					)
					RETURNING movie_id`

	var movieID int
	err := s.db.QueryRow(ctx, insertMovie, movie.Title, movie.Description, movie.Release, movie.Rating).Scan(&movieID)
	if err != nil {
		s.log.Error(err)
	}

	insertRelation := `INSERT INTO relations (actor_id, movie_id)
					   SELECT $1, $2
					   WHERE NOT EXISTS(
					       SELECT * FROM relations WHERE actor_id = $1::integer
					                               AND movie_id = $2::integer
					   )`

	for _, v := range actorsID {
		_, err := s.db.Exec(ctx, insertRelation, v, movieID)
		if err != nil {
			s.log.Error(err)
		}
	}

	return nil
}
