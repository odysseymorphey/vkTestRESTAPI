package server

import (
	"context"
	"github.com/odysseymorphey/vkTestRESTAPI/internal/dto"
)

type FilmaryService interface {
	CreateActor(ctx context.Context, actor dto.Actor) error
	UpdateActor(ctx context.Context, actor dto.Actor) error
	DeleteActor(ctx context.Context, id int) error
}
