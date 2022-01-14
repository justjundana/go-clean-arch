package news

import (
	_newsRepository "alta/cleanarch/businesses/news"
	"context"

	gorm "gorm.io/gorm"
)

type mysqlNewsRepository struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) _newsRepository.Repository {
	return &mysqlNewsRepository{
		Conn: conn,
	}
}

func (repository *mysqlNewsRepository) GetAll(ctx context.Context) ([]_newsRepository.Domain, error) {
	var data []News
	result := repository.Conn.Find(&data)
	if result.Error != nil {
		return []_newsRepository.Domain{}, result.Error
	}

	var domainNews []_newsRepository.Domain
	for _, value := range data {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, nil
}

func (repository *mysqlNewsRepository) GetByID(ctx context.Context, newsId int) (_newsRepository.Domain, error) {
	var data News
	result := repository.Conn.First(&data, newsId).Find(&data)
	if result.Error != nil {
		return _newsRepository.Domain{}, result.Error
	}

	return _newsRepository.Domain(data), nil
}
