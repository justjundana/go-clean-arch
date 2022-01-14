package news

import (
	"context"
	"time"
)

type newsUsecase struct {
	newsRepository Repository
	contextTimeout time.Duration
}

func NewNewsUsecase(nr Repository, timeout time.Duration) Usecase {
	return &newsUsecase{
		newsRepository: nr,
		contextTimeout: timeout,
	}
}

func (usecase *newsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	data, err := usecase.newsRepository.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (usecase *newsUsecase) GetByID(ctx context.Context, newsId int) (Domain, error) {
	data, err := usecase.newsRepository.GetByID(ctx, newsId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}
