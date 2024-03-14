package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Storage struct {
	log    *zap.SugaredLogger
	cancel context.CancelFunc
	db     *pgxpool.Pool
}

func NewStorage(log *zap.SugaredLogger, dsn string) (*Storage, error) {
	if log == nil {
		return nil, errors.New("Empty logger")
	}

	if dsn == "" {
		return nil, errors.New("Empty dsn")
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

	db.db = conn

	return db, nil
}

func (s *Storage) Close() {
	s.db.Close()
	s.cancel()
}

func (s *Storage) AddNewActor() {
	
}