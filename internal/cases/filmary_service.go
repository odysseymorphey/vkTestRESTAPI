package cases

import "go.uber.org/zap"

type Service struct {
	log *zap.SugaredLogger
	storage Storage
}
