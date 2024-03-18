package cases

import (
	"context"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/dto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type FilmaryService struct {
	log     *zap.SugaredLogger
	storage Storage
}

func NewFilmaryService(log *zap.SugaredLogger, storage Storage) (*FilmaryService, error) {
	if log == nil {
		return nil, errors.New("Empty logger")
	}

	if storage == nil || storage == Storage(nil) {
		return nil, errors.New("Empty storage")
	}

	return &FilmaryService{
		log:     log,
		storage: storage,
	}, nil
}

func (s *FilmaryService) CreateActor(ctx context.Context, actor dto.Actor) error {
	err := s.storage.CreateActor(ctx, actor)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *FilmaryService) UpdateActor(ctx context.Context, actor dto.Actor) error {
	err := s.storage.UpdateActor(ctx, actor)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *FilmaryService) DeleteActor(ctx context.Context, id int) error {
	err := s.storage.DeleteActor(ctx, id)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *FilmaryService) CreateMovie(ctx context.Context, movie dto.Movie) error {
	err := s.storage.CreateMovie(ctx, movie)
	if err != nil {
		s.log.Error(err)
	}

	return err
}

func (s *FilmaryService) DeleteMovie(ctx context.Context, movie dto.Movie) error {
	err := s.storage.DeleteMovie(ctx, movie)
	if err != nil {
		s.log.Error(err)
	}

	return err
}
