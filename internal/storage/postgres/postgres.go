package postgres

import (
	"context"
	"fmt"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
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
			  VALUES ($1, $2, $3)`

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
	panic(".")
}
